#![allow(non_snake_case)]
#![cfg_attr(
  all(not(debug_assertions), target_os = "windows"),
  windows_subsystem = "windows"
)]

use std::{collections::HashMap, fs, path::Path};

use chrono::Utc;
use notify_rust::Notification;
use reqwest::header::{self, HeaderMap};
use serde::{Deserialize, Serialize};
use serde_json::Value;
use tauri::{
  CustomMenuItem, Manager, SystemTray, SystemTrayEvent, SystemTrayMenu, SystemTrayMenuItem,
};

const CONFIG_FILE_NAME: &str = "config.json";

fn get_or_build_config_dir() -> String {
  let home_path = home::home_dir()
    .unwrap()
    .into_os_string()
    .into_string()
    .unwrap();

  let config_path = format!("{}{}", home_path, "/.config/galata/");

  let config_file_path = format!("{}{}", config_path, CONFIG_FILE_NAME);

  if !Path::new(&config_path).is_dir() {
    fs::create_dir(config_path).unwrap();
  }

  if !Path::new(&config_file_path).is_file() {
    fs::write(&config_file_path, "{}").unwrap();
  }

  return config_file_path;
}

fn get_auth_token() -> String {
  let config_file_path = get_or_build_config_dir();

  let config_as_string = fs::read_to_string(&config_file_path).unwrap();

  let config: HashMap<&str, Value> = serde_json::from_str(config_as_string.as_str()).unwrap();

  match config.get("elid") {
    Some(res) => return res.to_string().replace("\"", ""),
    None => return "".to_string(),
  }
}

fn get_refresh_token() -> String {
  let config_file_path = get_or_build_config_dir();

  let config_as_string = fs::read_to_string(&config_file_path).unwrap();

  let config: HashMap<&str, Value> = serde_json::from_str(config_as_string.as_str()).unwrap();

  match config.get("ucid") {
    Some(res) => return res.to_string().replace("\"", ""),
    None => return "".to_string(),
  }
}

fn set_auth_token(token: String) {
  let config_file_path = get_or_build_config_dir();

  let config_as_string = fs::read_to_string(&config_file_path).unwrap();

  let mut config: HashMap<&str, Value> = serde_json::from_str(config_as_string.as_str()).unwrap();

  config.insert("elid", serde_json::Value::String(token));

  let config_to_write = serde_json::to_string_pretty(&config).unwrap();
  fs::write(&config_file_path, config_to_write).unwrap();
}

fn set_refresh_token(token: String) {
  let config_file_path = get_or_build_config_dir();

  let config_as_string = fs::read_to_string(&config_file_path).unwrap();

  let mut config: HashMap<&str, Value> = serde_json::from_str(config_as_string.as_str()).unwrap();

  config.insert("ucid", serde_json::Value::String(token));

  let config_to_write = serde_json::to_string_pretty(&config).unwrap();
  fs::write(&config_file_path, config_to_write).unwrap();
}

fn create_cookie_headers() -> HeaderMap {
  let cookie = format!("ucid={}", get_refresh_token());
  let cookie_header = header::HeaderValue::from_str(cookie.as_str());
  let mut request_headers = header::HeaderMap::new();

  match cookie_header {
    Ok(header) => {
      request_headers.insert(header::COOKIE, header);
      return request_headers;
    }
    Err(_) => return request_headers,
  }
}

#[tokio::main]
async fn main() {
  let _ = get_or_build_config_dir();
  let dashboard = CustomMenuItem::new("dashboard".to_string(), "Dashboard");
  let download_update = CustomMenuItem::new("download_update".to_string(), "Download Update");
  let quit = CustomMenuItem::new("quit".to_string(), "Quit");
  let tray_menu = SystemTrayMenu::new()
    .add_item(dashboard)
    .add_item(download_update)
    .add_native_item(SystemTrayMenuItem::Separator)
    .add_item(quit);
  tauri::Builder::default()
    .invoke_handler(tauri::generate_handler![
      create_new_application,
      create_new_team,
      delete_application,
      delete_team,
      fetch_active_user,
      fetch_application_by_id,
      fetch_cached_alerts,
      fetch_auth_token,
      fetch_team_by_id,
      login,
      notify_user,
      register_user,
      remove_user_from_team,
    ])
    .system_tray(SystemTray::new().with_menu(tray_menu))
    .on_system_tray_event(|app, event| match event {
      SystemTrayEvent::MenuItemClick { id, .. } => {
        // get a handle to the clicked menu item
        // note that `tray_handle` can be called anywhere,
        // just get a `AppHandle` instance with `app.handle()` on the setup hook
        // and move it to another function or thread
        let window = app.get_window("main").unwrap();
        match id.as_str() {
          "dashboard" => {
            window.show().unwrap();
          }
          "download_update" => {
            // let window = app.get_window("main").unwrap();
            // window.emit("tauri://update", None);
            // Attempt to download new update. Display alert if one is not available.
          }
          "quit" => {
            window.close().unwrap();
          }
          _ => {}
        }
      }
      _ => {}
    })
    .on_window_event(|event| match event.event() {
      tauri::WindowEvent::CloseRequested { api, .. } => {
        event.window().hide().unwrap();
        api.prevent_close();
      }
      _ => {}
    })
    .run(tauri::generate_context!())
    .expect("error while building tauri application");
}

#[derive(Debug, Serialize, Deserialize)]
struct LoginPayload {
  email: String,
  password: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct LoginResponse {
  status: String,
  message: String,
  data: String,
}

#[tauri::command]
async fn login(email: String, password: String) -> Result<LoginResponse, String> {
  let client = reqwest::Client::builder()
    .cookie_store(true)
    .build()
    .unwrap();

  let url = "http://localhost:5000/api/login";

  let payload = LoginPayload {
    email: email,
    password: password,
  };

  let result = client
    .post(url)
    .json::<LoginPayload>(&payload)
    .send()
    .await
    .unwrap();

  let cookies_itr = result.cookies();

  for cookie in cookies_itr {
    if cookie.name() == "ucid" {
      set_refresh_token(cookie.value().to_string())
    }
  }

  let res = result.json::<LoginResponse>().await;

  match res {
    Ok(res) => {
      let token = res.data.to_owned();
      set_auth_token(token);
      Ok(res)
    }
    Err(e) => Err(format!("An error occurred {}", e)),
  }
}

#[derive(Debug, Serialize, Deserialize)]
struct RegisterUserPayload {
  first_name: String,
  last_name: String,
  email: String,
  password: String,
  password_confirmation: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct RegisterUserResponse {
  status: String,
  message: String,
  data: String,
}

#[tauri::command]
async fn register_user(
  register_user_input: RegisterUserPayload,
) -> Result<RegisterUserResponse, String> {
  let client = reqwest::Client::new();
  let url = "http://localhost:5000/api/register";

  let result = client
    .post(url)
    .json::<RegisterUserPayload>(&register_user_input)
    .send()
    .await
    .unwrap()
    .json::<RegisterUserResponse>()
    .await;

  match result {
    Ok(res) => {
      let token = res.data.to_owned();
      set_auth_token(token);
      Ok(res)
    }
    Err(e) => Err(format!("An error occurred {}", e)),
  }
}

#[derive(Debug, Serialize, Deserialize)]
struct FetchAuthTokenResponse {
  token: String,
}

#[tauri::command]
async fn fetch_auth_token() -> Result<FetchAuthTokenResponse, String> {
  let client = reqwest::Client::builder()
    .default_headers(create_cookie_headers())
    .cookie_store(true)
    .build()
    .unwrap();

  let url = "http://localhost:5000/api/refresh_token";

  let result = client.post(url).send().await.unwrap();

  let res = result.json::<FetchAuthTokenResponse>().await;

  match res {
    Ok(res) => {
      let token = res.token.to_owned();
      set_auth_token(token);
      Ok(res)
    }
    Err(e) => Err(format!("An error occurred {}", e)),
  }
}

#[derive(Debug, Serialize, Deserialize)]
struct Alert {
  Title: String,
  Description: String,
  Link: String,
}

#[derive(Debug, Serialize, Deserialize)]

struct CachedAlertsResponse {
  status: String,
  message: String,
  data: Option<Vec<Value>>,
}

#[tauri::command]
async fn fetch_cached_alerts(application_id: i32) -> Result<CachedAlertsResponse, String> {
  let client = reqwest::Client::new();
  let url = format!(
    "http://localhost:5000/api/applications/{}/alerts",
    application_id
  );
  let auth_token = get_auth_token();

  let result = client
    .get(url)
    .bearer_auth(auth_token)
    .send()
    .await
    .unwrap()
    .json::<CachedAlertsResponse>()
    .await;

  match result {
    Ok(res) => Ok(res),
    Err(err) => Err(format!(
      "An error occurred while fetching alerts {}",
      err.to_string()
    )),
  }
}

#[derive(Debug, Serialize, Deserialize)]
struct User {
  ID: i32,
  CreatedAt: chrono::DateTime<Utc>,
  UpdatedAt: Option<chrono::DateTime<Utc>>,
  DeletedAt: Option<chrono::DateTime<Utc>>,
  FirstName: String,
  LastName: String,
  Email: String,
  Password: String,
  IsAdmin: bool,
  IsVerified: bool,
  Applications: Vec<Application>,
  Teams: Option<Vec<Team>>,
}

#[derive(Debug, Serialize, Deserialize)]
struct FetchActiveUserResponse {
  status: String,
  message: String,
  data: Option<User>,
}

#[tauri::command]
async fn fetch_active_user() -> Result<FetchActiveUserResponse, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/users/me");
  let auth_token = get_auth_token();

  let result = client
    .get(url)
    .bearer_auth(auth_token)
    .send()
    .await
    .unwrap()
    .json::<FetchActiveUserResponse>()
    .await;

  match result {
    Ok(res) => Ok(res),
    Err(err) => Err(format!(
      "An error occurred while fetching active user {}",
      err.to_string()
    )),
  }
}

#[derive(Clone, Serialize, Deserialize)]
struct NewApplicationPayload {
  application_name: String,
  team_id: Option<i32>,
  user_id: Option<i32>,
  alert_schema: Option<AlertSchemaInput>,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
struct AlertSchemaInput {
  title: String,
  description: String,
  link: String,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
struct AlertSchema {
  Title: String,
  Description: String,
  Link: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct Application {
  ID: i32,
  Name: String,
  UniqueId: String,
  CreatedAt: chrono::DateTime<Utc>,
  UpdatedAt: Option<chrono::DateTime<Utc>>,
  DeletedAt: Option<chrono::DateTime<Utc>>,
  TeamId: Option<i32>,
  UserId: Option<i32>,
  AlertSchema: Option<AlertSchema>,
  User: Option<User>,
  Team: Option<Team>,
}

#[derive(Debug, Serialize, Deserialize)]
struct NewApplicationResponse {
  status: String,
  message: String,
  data: Application,
}

#[tauri::command]
async fn create_new_application(
  application_name: String,
  team_id: Option<i32>,
  user_id: Option<i32>,
  alert_schema: Option<AlertSchemaInput>,
) -> Result<NewApplicationResponse, String> {
  let client = reqwest::Client::new();
  let url = "http://localhost:5000/api/applications";
  let auth_token = get_auth_token();

  let mut payload = NewApplicationPayload {
    alert_schema: alert_schema,
    application_name: application_name,
    team_id: None,
    user_id: None,
  };

  // TODO: Add else clause that will throw an error if team_id or user_id is not found
  if team_id != None {
    payload.team_id = team_id;
  } else if user_id != None {
    payload.user_id = user_id;
  }

  let result = client
    .post(url)
    .bearer_auth(auth_token)
    .json::<NewApplicationPayload>(&payload)
    .send()
    .await
    .unwrap()
    .json::<NewApplicationResponse>()
    .await;

  match result {
    Ok(res) => Ok(res),
    Err(e) => Err(format!("An error occurred {}", e)),
  }
}

#[derive(Debug, Serialize, Deserialize)]
struct Team {
  ID: i32,
  Name: String,
  CreatedAt: chrono::DateTime<Utc>,
  UpdatedAt: Option<chrono::DateTime<Utc>>,
  DeletedAt: Option<chrono::DateTime<Utc>>,
  Applications: Option<Vec<Application>>,
  Managers: Option<Vec<User>>,
}

#[derive(Debug, Serialize, Deserialize)]
struct NewTeamPayload {
  team_name: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct NewTeamResponse {
  status: String,
  message: String,
  data: Team,
}

#[tauri::command]
async fn create_new_team(team_name: String) -> Result<NewTeamResponse, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/teams");
  let auth_token = get_auth_token();

  let payload = NewTeamPayload {
    team_name: team_name,
  };

  let result = client
    .post(url)
    .bearer_auth(auth_token)
    .json::<NewTeamPayload>(&payload)
    .send()
    .await
    .unwrap()
    .json()
    .await;

  match result {
    Ok(res) => Ok(res),
    Err(e) => Err(format!("An error occurred {}", e)),
  }
}

#[derive(Debug, Serialize, Deserialize)]
struct FetchTeamByIdPayload {
  status: String,
  message: String,
  data: Team,
}

#[tauri::command]
async fn fetch_team_by_id(team_id: i32) -> Result<String, ()> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/teams/{}", team_id);
  let auth_token = get_auth_token();

  let result = client
    .get(url)
    .bearer_auth(auth_token)
    .send()
    .await
    .unwrap()
    .text()
    .await
    .unwrap();

  Ok(result)
}

#[derive(Debug, Serialize, Deserialize)]
struct DeleteTeamPayload {
  status: String,
  message: String,
}

#[tauri::command]
async fn delete_team(team_id: i32) -> Result<DeleteTeamPayload, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/teams/{}", team_id);
  let auth_token = get_auth_token();

  let result = client
    .delete(url)
    .bearer_auth(auth_token)
    .send()
    .await
    .unwrap()
    .json::<DeleteTeamPayload>()
    .await;

  match result {
    Ok(res) => Ok(res),
    Err(err) => Err(format!(
      "An error occurred while fetching application {}",
      err.to_string()
    )),
  }
}

#[tauri::command]
async fn remove_user_from_team(team_id: i32, user_id: i32) -> Result<String, ()> {
  let client = reqwest::Client::new();
  let url = format!(
    "http://localhost:5000/api/teams/{}/user/{}",
    team_id, user_id
  );
  let auth_token = get_auth_token();

  let result = client
    .delete(url)
    .bearer_auth(auth_token)
    .send()
    .await
    .unwrap()
    .text()
    .await
    .unwrap();

  Ok(result)
}

#[derive(Debug, Serialize, Deserialize)]
struct FetchApplicationByIdPayload {
  status: String,
  message: String,
  data: Application,
}

#[tauri::command]
async fn fetch_application_by_id(
  application_id: i32,
) -> Result<FetchApplicationByIdPayload, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/applications/{}", application_id);
  let auth_token = get_auth_token();

  let result = client
    .get(url)
    .bearer_auth(auth_token)
    .send()
    .await
    .unwrap()
    .json::<FetchApplicationByIdPayload>()
    .await;

  match result {
    Ok(res) => Ok(res),
    Err(err) => Err(format!(
      "An error occurred while fetching application {}",
      err.to_string()
    )),
  }
}

#[derive(Debug, Serialize, Deserialize)]
struct DeleteApplicationPayload {
  status: String,
  message: String,
}

#[tauri::command]
async fn delete_application(application_id: i32) -> Result<DeleteApplicationPayload, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/applications/{}", application_id);
  let auth_token = get_auth_token();

  let result = client
    .delete(url)
    .bearer_auth(auth_token)
    .send()
    .await
    .unwrap()
    .json::<DeleteApplicationPayload>()
    .await;

  match result {
    Ok(res) => Ok(res),
    Err(err) => Err(format!(
      "An error occurred while fetching application {}",
      err.to_string()
    )),
  }
}

#[tauri::command]
fn notify_user(title: &str, body: &str) {
  Notification::new()
    .summary(title)
    .body(body)
    .show()
    .unwrap();
}

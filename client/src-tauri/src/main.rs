#![allow(non_snake_case)]
#![cfg_attr(
  all(not(debug_assertions), target_os = "windows"),
  windows_subsystem = "windows"
)]

use chrono::Utc;
use notify_rust::Notification;
use serde::{Deserialize, Serialize};
use serde_json::Value;
use tauri::{
  CustomMenuItem, Manager, SystemTray, SystemTrayEvent, SystemTrayMenu, SystemTrayMenuItem,
};

#[tokio::main]
async fn main() {
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
      fetch_team_by_id,
      notify_user
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
async fn fetch_cached_alerts(
  application_id: i32,
  auth_token: &str,
) -> Result<CachedAlertsResponse, String> {
  let client = reqwest::Client::new();
  let url = format!(
    "http://localhost:5000/api/applications/{}/alerts",
    application_id
  );

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

#[tauri::command]
async fn fetch_active_user(auth_token: &str) -> Result<String, ()> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/users/me");

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
}

#[derive(Debug, Serialize, Deserialize)]
struct NewApplicationResponse {
  status: String,
  message: String,
  data: Application,
}

#[tauri::command]
async fn create_new_application(
  auth_token: String,
  application_name: String,
  team_id: Option<i32>,
  user_id: Option<i32>,
  alert_schema: Option<AlertSchemaInput>,
) -> Result<NewApplicationResponse, String> {
  let client = reqwest::Client::new();
  let url = "http://localhost:5000/api/applications";

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
  CreatedAt: chrono::DateTime<Utc>,
  UpdatedAt: Option<chrono::DateTime<Utc>>,
  DeletedAt: Option<chrono::DateTime<Utc>>,
  Name: String,
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
async fn create_new_team(auth_token: String, team_name: String) -> Result<NewTeamResponse, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/teams");

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
struct FetchTeamByPayload {
  status: String,
  message: String,
  data: Team,
}

#[tauri::command]
async fn fetch_team_by_id(auth_token: &str, team_id: i32) -> Result<FetchTeamByPayload, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/teams/{}", team_id);

  let result = client
    .get(url)
    .bearer_auth(auth_token)
    .send()
    .await
    .unwrap()
    .json::<FetchTeamByPayload>()
    .await;

  match result {
    Ok(res) => Ok(res),
    Err(err) => Err(format!(
      "An error occurred while fetching team {}",
      err.to_string()
    )),
  }
}

#[derive(Debug, Serialize, Deserialize)]
struct DeleteTeamPayload {
  status: String,
  message: String,
}

#[tauri::command]
async fn delete_team(auth_token: &str, team_id: i32) -> Result<DeleteTeamPayload, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/teams/{}", team_id);

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

#[derive(Debug, Serialize, Deserialize)]
struct FetchApplicationByPayload {
  status: String,
  message: String,
  data: Application,
}

#[tauri::command]
async fn fetch_application_by_id(
  auth_token: &str,
  application_id: i32,
) -> Result<FetchApplicationByPayload, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/applications/{}", application_id);

  let result = client
    .get(url)
    .bearer_auth(auth_token)
    .send()
    .await
    .unwrap()
    .json::<FetchApplicationByPayload>()
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
async fn delete_application(
  auth_token: &str,
  application_id: i32,
) -> Result<DeleteApplicationPayload, String> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/applications/{}", application_id);

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

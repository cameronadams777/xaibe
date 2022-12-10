#![allow(non_snake_case)]
#![cfg_attr(
  all(not(debug_assertions), target_os = "windows"),
  windows_subsystem = "windows"
)]

use std::{collections::HashMap, env, fs, path::Path};

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
      get_stored_auth_token,
      logout_user,
      notify_user,
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

#[tauri::command]
fn get_stored_auth_token() -> String {
  return get_auth_token();
}

#[tauri::command]
fn logout_user() {
  set_auth_token("".to_string());
  set_refresh_token("".to_string());
}

#[tauri::command]
fn notify_user(title: &str, body: &str) {
  Notification::new()
    .summary(title)
    .body(body)
    .show()
    .unwrap();
}


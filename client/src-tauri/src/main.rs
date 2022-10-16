#![cfg_attr(
  all(not(debug_assertions), target_os = "windows"),
  windows_subsystem = "windows"
)]

use serde_json::Value;

#[tokio::main]
async fn main() {
  tauri::Builder::default()
    .invoke_handler(tauri::generate_handler![
      fetch_cached_alerts,
      fetch_active_user,
    ])
    .run(tauri::generate_context!())
    .expect("error while running tauri application");
}

#[tauri::command]
async fn fetch_cached_alerts(
  application_id: i32,
  auth_token: &str,
  service_token: &str,
) -> Result<String, ()> {
  let client = reqwest::Client::new();
  let url = format!(
    "http://localhost:5000/api/applications/{}/alerts",
    application_id
  );
  let bearer_token = format!("Bearer {}", auth_token);

  let result = client
    .get(url)
    .header("Authorization", bearer_token)
    .send()
    .await
    .unwrap()
    .text()
    .await
    .unwrap();

  Ok(result)
}

#[tauri::command]
async fn fetch_active_user(auth_token: &str) -> Result<String, ()> {
  let client = reqwest::Client::new();
  let url = format!("http://localhost:5000/api/users/me");
  let bearer_token = format!("Bearer {}", auth_token);

  let result = client
    .get(url)
    .header("Authorization", bearer_token)
    .send()
    .await
    .unwrap()
    .text()
    .await
    .unwrap();

  Ok(result)
}

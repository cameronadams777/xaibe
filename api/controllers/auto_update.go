package controllers

import (
  "api/structs"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PlatformResponse struct {
  Signature string `json:"signature"`
  URL string `json:"url"`
}

type XaibeReleaseResponse struct {
  Version string `json:"version"`
  Notes string `json:"notes"`
  PubDate string `json:"pub_date"`
  Platforms map[string]PlatformResponse `json:"platforms"`
}

type GitHubReleaseAsset struct {
  URL string `json:"url"`
  ID int `json:"id"`
  Name string `json:"name"`
  NodeID string `json:"node_id"`
  Label *string `json:"label"`
  ContentType string `json:"content_type"`
  State string `json:"state"`
  Size int `json:"size"`
  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
  BrowserDownloadURL string `json:"browser_download_url"`
}

type GitHubResponse struct {
  TagName string `json:"tag_name"`
  PublishedAt string `json:"published_at"`
  Body string `json:"body"`
  Assets []GitHubReleaseAsset `json:"assets"`
}

var PLATFORMS = []string{
  "darwin-x86_64",
  "darwin-aarch64",
  "windows-x86_64",
}

const XAIBE_DESKTOP_RELEASES_REPO = "Galata-App/xaibe-releases"
  
func get_latest_gh_release(repo string, requested_platform string) (*XaibeReleaseResponse, error) {
  github_latest_release_url := "https://api.github.com/repos/" + repo + "/releases/latest"; 
  
  var github_response GitHubResponse

  resp, err := http.Get(github_latest_release_url)

  if err != nil {
    return nil, err 
  }

  defer resp.Body.Close()

  if decode_err := json.NewDecoder(resp.Body).Decode(&github_response); decode_err != nil {
     log.Fatal(decode_err)
     return nil, decode_err 
  }

  release_response := XaibeReleaseResponse{
    Version: github_response.TagName,
    Notes: github_response.Body,
    PubDate: github_response.PublishedAt,
    Platforms: make(map[string]PlatformResponse),
  }

  for _, asset := range github_response.Assets {
    if strings.Contains(asset.Name, requested_platform) {
      release_response.Platforms[requested_platform] = PlatformResponse{
        Signature: "",
        URL: asset.BrowserDownloadURL,
      }
    }
  }

  if len(release_response.Platforms) == 0 {
    return nil, errors.New("Xaibe Error: Requested platform not supported")
  }

  return &release_response, nil
}

func CheckLatestAppVersion(c *gin.Context) {
	current_version := c.Param("current_version")
  requested_platform := c.Param("platform")
  
  latest_release, fetch_rel_err := get_latest_gh_release(XAIBE_DESKTOP_RELEASES_REPO, requested_platform)

  if fetch_rel_err != nil {
    log.Fatal(fetch_rel_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: ""})
    return
  }

  latest_version := latest_release.Version
  formatted_version := strings.Replace(latest_version, "v", "", -1)
  version_types := strings.Split(formatted_version, ".")
  latest_maj := version_types[0]
  latest_min := version_types[1]
  latest_patch := version_types[2]

  current_version_type := strings.Split(current_version, ".")
  current_maj := current_version_type[0]
  current_min := current_version_type[1]
  current_patch := current_version_type[2]
  
  if current_maj == latest_maj && current_min == latest_min && current_patch== latest_patch {
    c.JSON(http.StatusNoContent, "")
    return
  }
  
  c.JSON(http.StatusOK, latest_release)
}

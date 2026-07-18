use serde::Deserialize;
use std::path::Path;

#[derive(Debug, Deserialize, Clone)]
pub struct PackageManifest {
    pub package: PackageInfo,
    pub dependencies: Option<Vec<String>>,
    pub install: InstallConfig,
}

#[derive(Debug, Deserialize, Clone)]
pub struct PackageInfo {
    pub name: String,
    pub version: String,
    pub description: String,
    pub homepage: Option<String>,
    pub repository: Option<String>,
}

#[derive(Debug, Deserialize, Clone)]
pub struct InstallConfig {
    #[serde(rename = "type")]
    pub install_type: String,
    pub url: Option<String>,
    pub strip_prefix: Option<String>,
    pub build_cmd: Option<String>,
    pub bin: Option<Vec<String>>,
}

impl PackageManifest {
    pub fn from_file(path: &Path) -> anyhow::Result<Self> {
        let content = std::fs::read_to_string(path)?;
        let manifest: PackageManifest = toml::from_str(&content)?;
        Ok(manifest)
    }

    pub fn render_url(&self) -> Option<String> {
        let url = self.install.url.as_ref()?;
        let version = &self.package.version;
        let name = &self.package.name;
        let rendered = url
            .replace("{{version}}", version)
            .replace("{{name}}", name);
        Some(rendered)
    }
}

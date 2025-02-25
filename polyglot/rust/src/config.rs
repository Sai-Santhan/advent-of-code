use std::path::PathBuf;
use anyhow::{anyhow, Context, Result};
use crate::opts::Opts;

#[derive(Debug)]
pub struct Config {
    pub operation: Operation,
    pub pwd: PathBuf,
    pub config: PathBuf,
}

impl TryFrom<Opts> for Config {
    type Error = anyhow::Error;

    fn try_from(value: Opts) -> Result<Self> {
        let operation = value.args.try_into()?;
        let config = get_config(value.config)?;
        let pwd = get_pwd(value.pwd)?;

        Ok(Config {
            operation,
            pwd,
            config,
        })
    }
}

#[derive(Debug, PartialEq)]
pub enum Operation {
    Print(Option<String>),
    Add(String, String),
    Remove(String),
}

impl TryFrom<Vec<String>> for Operation {
    type Error = anyhow::Error;

    fn try_from(value: Vec<String>) -> Result<Self, Self::Error> {
        let mut value = value;

        if value.is_empty() {
            return Ok(Operation::Print(None));
        }

        let term = value.get(0).expect("Expected to exist.");

        if term == "add" {
            if value.len() != 3 {
                return Err(anyhow!("Operation add expects 2 arguments but got {}", value.len()-1));
            }
            let mut drain = value.drain(1..=2);
            return Ok(Operation::Add(
                drain.next().expect("to exist"),
                drain.next().expect("to exist"),
            ));
        }

        if term == "remove" {
            if value.len() != 2 {
                return Err(anyhow!("Operation remove expects 1 arguments but got {}", value.len()-1));
            }
            let arg = value.pop().expect("to exist");
            return Ok(Operation::Remove(arg));
        }

        if value.len() > 1 {
            return Err(anyhow!("Operation print expects 0 or 1 arguments but got {}", value.len()));
        }

        let arg = value.pop().expect("to exist");
        Ok(Operation::Print(Some(arg)))
    }
}

fn get_config(config: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(v) = config {
        return Ok(v);
    }

    let loc = std::env::var("XDG_CONFIG_HOME")
        .context("Unable to get XDG_CONFIG_HOME")?;

    let mut loc = PathBuf::from(loc);
    loc.push("projector");
    loc.push("projector.json");
    Ok(loc)
}

fn get_pwd(pwd: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(pwd) = pwd {
        return Ok(pwd);
    }
    std::env::current_dir().context("Unable to get CURRENT_DIR")
}

#[cfg(test)]
mod test {
    use crate::opts::Opts;
    use anyhow::Result;
    use crate::config::Operation;
    use super::Config;

    #[test]
    fn test_print_all() -> Result<()> {
        let opts: Config = Opts {
            args: vec![],
            pwd: None,
            config: None,
        }.try_into()?;

        assert_eq!(opts.operation, Operation::Print(None));
        Ok(())
    }

    #[test]
    fn test_print_key() -> Result<()> {
        let opts: Config = Opts {
            args: vec!["foo".to_string()],
            pwd: None,
            config: None,
        }.try_into()?;

        assert_eq!(opts.operation, Operation::Print(Some("foo".to_string())));
        Ok(())
    }

    #[test]
    fn test_add_key_value() -> Result<()> {
        let opts: Config = Opts {
            args: vec![
                "add".to_string(),
                "foo".to_string(),
                "bar".to_string(),
            ],
            pwd: None,
            config: None,
        }.try_into()?;

        assert_eq!(opts.operation, Operation::Add(
            "foo".to_string(),
            "bar".to_string()));
        Ok(())
    }

    #[test]
    fn test_remove_key() -> Result<()> {
        let opts: Config = Opts {
            args: vec![
                "add".to_string(),
                "foo".to_string(),
            ],
            pwd: None,
            config: None,
        }.try_into()?;

        assert_eq!(opts.operation, Operation::Remove(
            "foo".to_string()));
        Ok(())
    }
}
# Changelog

## v0.1.0 - CLI & Configuration

- Add configuration with Viper
- Enhance CLI usage with Cobra
- Add Version command
- Add Config command (to print the config)
- Add Completion command (to generate completion data for shells), based on cobra's template.
- Enhance build process
  - Links version and build informations into binaries.
- Moved most packages to /internal/
- Rewrited some code for readability

## v0.0.0 - Init

Initial version.

- Add server command with:
  - Redirection feature.
    - stee can now send http redirections to clients based the path they requested and an embedded KV store (file storage).
  - Simple API to CRUD the redirections.

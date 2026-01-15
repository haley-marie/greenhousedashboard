# Greenhouse Dashboard (Desktop, Local, Offline-First)

## Overview
A lightweight, offline first desktop application for small growers and local farmers to track plantings, care events, and greenhouse conditions without requiring reliable internet access.

## Problem
Many small growers operate in environments with unreliable or nonexistent internet connectivity. Existing farm management tools are often cloud-dependent, overly complex on the user-end, or designed for large-scale operations. 

With increasing privacy and security issues, software with a small attack surface will be invaluable for small growers and farmers who cannot afford a data breach and do not have the technical experience necessary to safeguard themselves alone.

This project explores an alternative: a local-first, accessible application designed for real-world constraints and users.

## Design Goals
- Works fully offline
- Accessible to non-technical users
- Lightweight and fast on low-end hardware
- Focused on small farms and greenhouses
- Minimal setup and maintenance
- A smooth user experience. "I installed this and it just works"

## Architecture
- Frontend: HTML, CSS, and vanilla JavaScript
- Desktop shell: Tauri
- Backend logic: Go
- Database: SQLite (local file)

## Why Offline-First?
This application prioritizes local data storage and usability in rural environments. All data is stored locally and remains usable during power outages or connectivity loss.

## Data Model (High-Level)
- Growing areas (beds, greenhouses, fields)
- Plantings (groups of plants by area)
- Care events (watering, fertilizing, harvesting)
- Optional schedules and environmental readings

## Current Status
This project is under active development. Early focus is on:
- Core data model
- Greenhouse dashboard UI
- Local persistence and migrations

## Future Directions
- Optional data export
- Sensor integration
- Accessibility improvements
- Multi-greenhouse support

## Notes
This project is part of an ongoing exploration of product design, UX, and software architecture for sustainable agriculture. I believe in software made to reduce the human cognitive load of mundane tasks rather than to replace our hands-on interaction with our work.

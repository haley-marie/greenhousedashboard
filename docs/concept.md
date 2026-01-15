# Offline-First Smart Greenhouse Dashboard for Small Growers  
- standalone app over web  
- offline-first over cloud-dependent  
- non technical user over developer  
- local data storage over SaaS lock in  

## Architecture  
### Backend:  
- Go  
- SQLite  
### Frontend:  
- Tauri  
- HTML/CSS/JS  

## DB Structure  
### plants  
- id INTEGER PRIMARY KEY  
- species  
- nickname (optional)  
- variety (optional)  
- created_at  
- updated_at  

### growing_areas  
- id INTEGER PRIMARY KEY  
- name  
- type (greenhouse, bed, field, indoor)  
- location_notes  
- created_at  
- updated_at  

### plantings  
- id INTEGER PRIMARY KEY  
- plant_id  
- growing_area_id  
- quantity  
- planted_at  
- expected_harvest_at  
- status (active, harvested, failed)  

### care_events  
- id INTEGER PRIMARY KEY  
- planting_id OR growing_area_id  
- event_type (water, fertilize, prune, harvest)  
- notes  
- occurred_at  
- created_at  

### care_schedules  
- id INTEGER PRIMARY KEY  
- applies_to (planting | growing_area)  
- event_type  
- frequency (daily, weekly, custom)  
- start_date  
- end_date  

### environment_readings  
- id INTEGER PRIMARY KEY  
- growing_area_id  
- temperature  
- humidity  
- soil_moisture  
- recorded_at  


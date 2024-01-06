### CRUD APP GOLANG
Dependencies injected in Project
- DB => Repository
- Repository => Controller
- Controller => Handler

### Reference about table creation
```sql```
CREATE TABLE IF NOT EXISTS vendors
(
    id                          BIGINT PRIMARY KEY,
    plant                       TEXT                                   NOT NULL CONSTRAINT vendors_plant_check CHECK (plant <> ''::TEXT),
    status                      TEXT                                   NOT NULL CONSTRAINT vendors_status_check CHECK (status <> ''::TEXT),
    vendor_no                   TEXT                                   NOT NULL CONSTRAINT vendors_vendor_no_check CHECK (vendor_no <> ''::TEXT),
    vendor_name                 TEXT,
    storage_location            TEXT,
    company_code                TEXT,
    purch_org                   TEXT,
    uen_no                      TEXT,
    email                       TEXT,
    address                     TEXT,
    phone_number                TEXT,
    contact_person_firstname    TEXT,
    contact_person_lastname     TEXT,
    sap_is_posting_blocked      BOOLEAN                  DEFAULT FALSE NOT NULL,
    sap_is_purchasing_blocked   BOOLEAN                  DEFAULT FALSE NOT NULL,
    sap_is_flins_deleted        BOOLEAN                  DEFAULT FALSE NOT NULL,
    sap_is_vmi_confirmed        BOOLEAN                  DEFAULT FALSE NOT NULL,
    sap_is_vmi_temp_blocked     BOOLEAN                  DEFAULT FALSE NOT NULL,
    created_at                  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at                  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    UNIQUE (vendor_no)
);
CREATE INDEX IF NOT EXISTS vendors_vendor_no_index ON vendors(vendor_no);
CREATE INDEX IF NOT EXISTS vendors_storage_location_index ON vendors(storage_location);
``````

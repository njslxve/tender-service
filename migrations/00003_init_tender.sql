-- +goose Up
CREATE TABLE IF NOT EXISTS tenders (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  latest_version INT NOT NULL DEFAULT 1 CHECK(latest_version > 0)
);

CREATE TABLE IF NOT EXISTS tenders_versions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  tender_id UUID NOT NULL,
  name VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  service_type VARCHAR(100) NOT NULL,
  status VARCHAR(100) NOT NULL,
  creator_username VARCHAR(50) NOT NULL,
  organization_id UUID NOT NULL,
  version INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(tender_id) REFERENCES tenders(id) ON DELETE CASCADE,
  UNIQUE (tender_id, version)
);

-- +goose Down
DROP TABLE tenders;
DROP TABLE tenders_versions;
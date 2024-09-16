-- +goose Up
CREATE TABLE IF NOT EXISTS bids (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  latest_version INT NOT NULL DEFAULT 1 CHECK(latest_version > 0)
);

CREATE TABLE IF NOT EXISTS bids_versions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  bid_id UUID NOT NULL,
  name VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  tender_id UUID NOT NULL,
  status VARCHAR(100) NOT NULL,
  author_type VARCHAR(100) NOT NULL,
  author_id UUID NOT NULL,
  version INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(bid_id) REFERENCES bids(id) ON DELETE CASCADE,
  UNIQUE (bid_id, version)
);

CREATE TABLE IF NOT EXISTS bid_feedback (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  bid_id UUID NOT NULL,
  description TEXT NOT NULL,
  author_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(bid_id) REFERENCES bids(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS bids_decisions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  tender_id UUID NOT NULL,
  bid_id UUID NOT NULL,
  decision VARCHAR(100),
  approved_count INT NOT NULL DEFAULT 0,
  rejected_count INT NOT NULL DEFAULT 0,
  FOREIGN KEY(tender_id) REFERENCES tenders(id) ON DELETE CASCADE,
  FOREIGN KEY(bid_id) REFERENCES bids(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE bids;
DROP TABLE bids_versions;
DROP TABLE bid_feedback;
DROP TABLE bids_decisions;
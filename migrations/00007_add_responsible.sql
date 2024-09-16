-- +goose Up
INSERT INTO organization_responsible (organization_id, user_id)
VALUES ('2599da85-8a05-4c2f-bd4a-755c21cd788e', '7189d52d-bfcb-415b-9c7a-e5c16d5cae61'),
  ('d976cd81-3c1f-4d75-9841-1003af7d1e40', 'e430b712-acce-444a-977e-dc109e952420'),
  ('d976cd81-3c1f-4d75-9841-1003af7d1e40', 'ef958565-260c-452f-a28c-6c325366f3c9'),
  ('d976cd81-3c1f-4d75-9841-1003af7d1e40', '1c1ceef4-fd69-4f63-ac4b-f1cdf6568197'),
  ('d976cd81-3c1f-4d75-9841-1003af7d1e40', '92a911cc-4061-40bb-bd7a-48f7d7efe330');

-- +goose Down
DELETE FROM organization_responsible WHERE organization_id IN ('2599da85-8a05-4c2f-bd4a-755c21cd788e', 'd976cd81-3c1f-4d75-9841-1003af7d1e40');
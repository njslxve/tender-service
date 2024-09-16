-- +goose Up
INSERT INTO organization (id, name, description, type)
VALUES ('2599da85-8a05-4c2f-bd4a-755c21cd788e', 'Comodoro', 'Comodoro Brigade', 'LLC'),
  ('d976cd81-3c1f-4d75-9841-1003af7d1e40', 'Imperium', 'Imperium Team', 'IE');

-- +goose Down
DELETE FROM organisation WHERE name IN ('Comodoro', 'Imperium');
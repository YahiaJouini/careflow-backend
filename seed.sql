INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES (
  'Youssef',
  'Ben Ali',
  'youssef.benali@careflow.com',
  '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
  'https://images.unsplash.com/photo-1606813907291-d86efa9b94db?w=400&h=400&fit=crop',
  true,
  'doctor',
  NOW(),
  NOW(),
  NOW()
);

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT
  u.id,
  s.id,
  'Interventional cardiologist with a focus on coronary artery disease.',
  'CARD-TN-2001',
  130.00,
  true,
  true,
  NOW(),
  NOW()
FROM users u, specialties s
WHERE u.email = 'youssef.benali@careflow.com'
  AND s.name = 'Cardiologist';


INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES (
  'Amira',
  'Trabelsi',
  'amira.trabelsi@careflow.com',
  '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
  'https://images.unsplash.com/photo-1582750433449-648ed127bb54?w=400&h=400&fit=crop',
  true,
  'doctor',
  NOW(),
  NOW(),
  NOW()
);

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT
  u.id,
  s.id,
  'Dermatology specialist treating acne, eczema, and skin cancers.',
  'DERM-TN-3112',
  90.00,
  true,
  true,
  NOW(),
  NOW()
FROM users u, specialties s
WHERE u.email = 'amira.trabelsi@careflow.com'
  AND s.name = 'Dermatologist';


INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES (
  'Hichem',
  'Mansour',
  'hichem.mansour@careflow.com',
  '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
  'https://images.unsplash.com/photo-1550831107-1553da8c8464?w=400&h=400&fit=crop',
  true,
  'doctor',
  NOW(),
  NOW(),
  NOW()
);

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT
  u.id,
  s.id,
  'Orthopedic surgeon specializing in sports injuries and joint replacement.',
  'ORTHO-TN-4420',
  140.00,
  true,
  true,
  NOW(),
  NOW()
FROM users u, specialties s
WHERE u.email = 'hichem.mansour@careflow.com'
  AND s.name = 'Orthopedic';


INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES (
  'Salma',
  'Gharbi',
  'salma.gharbi@careflow.com',
  '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
  'https://images.unsplash.com/photo-1614608682850-e0d6ed316d47?w=400&h=400&fit=crop',
  true,
  'doctor',
  NOW(),
  NOW(),
  NOW()
);

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT
  u.id,
  s.id,
  'Compassionate pediatrician focused on early childhood development.',
  'PED-TN-1789',
  70.00,
  true,
  true,
  NOW(),
  NOW()
FROM users u, specialties s
WHERE u.email = 'salma.gharbi@careflow.com'
  AND s.name = 'Pediatrician';


INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES (
  'Karim',
  'Bouazizi',
  'karim.bouazizi@careflow.com',
  '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
  'https://images.unsplash.com/photo-1607746882042-944635dfe10e?w=400&h=400&fit=crop',
  true,
  'doctor',
  NOW(),
  NOW(),
  NOW()
);

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT
  u.id,
  s.id,
  'Neurologist treating epilepsy, Parkinson’s disease, and stroke patients.',
  'NEURO-TN-5533',
  155.00,
  true,
  true,
  NOW(),
  NOW()
FROM users u, specialties s
WHERE u.email = 'karim.bouazizi@careflow.com'
  AND s.name = 'Neurologist';


INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES (
  'Noura',
  'Ayadi',
  'noura.ayadi@careflow.com',
  '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
  'https://images.unsplash.com/photo-1598257006463-7c64a5a538cc?w=400&h=400&fit=crop',
  true,
  'doctor',
  NOW(),
  NOW(),
  NOW()
);

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT
  u.id,
  s.id,
  'Ophthalmologist providing cataract and laser vision correction.',
  'OPH-TN-9011',
  95.00,
  true,
  true,
  NOW(),
  NOW()
FROM users u, specialties s
WHERE u.email = 'noura.ayadi@careflow.com'
  AND s.name = 'Ophthalmologist';

INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES ('Richard', 'Lee', 'richard.lee@careflow.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'https://images.unsplash.com/photo-1612349317150-e413f6a5b16d?auto=format&fit=crop&w=400&h=400', true, 'doctor', NOW(), NOW(), NOW());

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT 
    u.id, 
    s.id, 
    'Experienced dentist focusing on comprehensive oral hygiene and cosmetic procedures.', 
    'DENT-1001', 
    100.00, 
    true,       
    true, 
    NOW(), 
    NOW()
FROM users u, specialties s
WHERE u.email = 'richard.lee@careflow.com' AND s.name = 'Dentist';

INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES ('Sarah', 'Smith', 'sarah.smith@careflow.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'https://images.unsplash.com/photo-1559839734-2b71ea197ec2?q=80&w=1170&auto=format&fit=crop&w=400&h=400', true, 'doctor', NOW(), NOW(), NOW());

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT 
    u.id, 
    s.id, 
    'Expert in heart rhythm disorders and preventative cardiology with 15 years of experience.', 
    'CARD-8821', 
    120.00, 
    true, 
    true, 
    NOW(), 
    NOW()
FROM users u, specialties s
WHERE u.email = 'sarah.smith@careflow.com' AND s.name = 'Cardiologist';

INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES ('James', 'Chen', 'james.chen@careflow.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'https://images.unsplash.com/photo-1537368910025-700350fe46c7?auto=format&fit=crop&w=400&h=400', true, 'doctor', NOW(), NOW(), NOW());

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT 
    u.id, 
    s.id, 
    'Specialist in treating migraines, epilepsy, and other nervous system disorders.', 
    'NEU-4412', 
    150.00, 
    true, 
    true, 
    NOW(), 
    NOW()
FROM users u, specialties s
WHERE u.email = 'james.chen@careflow.com' AND s.name = 'Neurologist';

INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES ('Emily', 'Davis', 'emily.davis@careflow.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'https://images.unsplash.com/photo-1594824476967-48c8b964273f?auto=format&fit=crop&w=400&h=400', true, 'doctor', NOW(), NOW(), NOW());

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT 
    u.id, 
    s.id, 
    'Cosmetic and restorative dentistry specialist. Creating smiles that last a lifetime.', 
    'DENT-5541', 
    80.00, 
    true, 
    true, 
    NOW(), 
    NOW()
FROM users u, specialties s
WHERE u.email = 'emily.davis@careflow.com' AND s.name = 'Dentist';

INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES ('Michael', 'Brown', 'michael.brown@careflow.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'https://images.unsplash.com/photo-1622253692010-333f2da6031d?auto=format&fit=crop&w=400&h=400', true, 'doctor', NOW(), NOW(), NOW());

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT 
    u.id, 
    s.id, 
    'Providing advanced eye care and vision correction surgeries.', 
    'OPH-9932', 
    90.00, 
    true, 
    true, 
    NOW(), 
    NOW()
FROM users u, specialties s
WHERE u.email = 'michael.brown@careflow.com' AND s.name = 'Ophthalmologist';

INSERT INTO users (first_name, last_name, email, password, image, verified, role, code_expiration_time, created_at, updated_at)
VALUES ('Linda', 'Wilson', 'linda.wilson@careflow.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'https://plus.unsplash.com/premium_photo-1673953510158-174d4060db8b?q=80&w=687&auto=format&fit=crop&w=400&h=400', true, 'doctor', NOW(), NOW(), NOW());

INSERT INTO doctors (user_id, specialty_id, bio, license_number, consultation_fee, is_available, is_verified, created_at, updated_at)
SELECT 
    u.id, 
    s.id, 
    'Dedicated to women''s heart health and preventive cardiology.', 
    'CARD-7731', 
    115.00, 
    true, 
    true, 
    NOW(), 
    NOW()
FROM users u, specialties s
WHERE u.email = 'linda.wilson@careflow.com' AND s.name = 'Cardiologist';


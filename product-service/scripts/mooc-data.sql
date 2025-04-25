INSERT INTO public.categories (id, name, description, parent_id, image_url, is_active, created_at, updated_at, deleted_at)
VALUES
(1, 'Coques et étuis', 'Accessoires pour la protection des smartphones', NULL, 'https://www.laredoute.fr/ppdp/prod-601377272.aspx', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),
(2, 'Écouteurs', 'Écouteurs filaires et sans fil pour smartphones', NULL, 'https://www.pexels.com/fr-fr/photo/oreillettes-blanches-avec-fil-373945/', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),
(3, 'Chargeurs et câbles', 'Chargeurs rapides, câbles USB, et stations de charge', NULL, 'https://www.pexels.com/fr-fr/photo/chargeur-de-telephone-portable-blanc-sur-fond-blanc-1037992/', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),
(4, 'Protections d''écran', 'Films en verre trempé, protections d''écran anti-rayures', NULL, 'https://www.pexels.com/fr-fr/photo/smartphone-noir-avec-protection-d-ecran-267394/', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),
(5, 'Supports et docks', 'Supports voiture, trépieds, et stations de recharge sans fil', NULL, 'https://www.pexels.com/fr-fr/photo/support-de-telephone-pour-voiture-386009/', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL);

INSERT INTO public.products (
  id, name, description, price, sku, stock_level, image_url, category_id, attributes,
  is_active, created_at, updated_at, deleted_at
)
VALUES
-- Coques et étuis
(1, 'Coque Silicone iPhone 13', 'Coque souple en silicone, antichoc, pour iPhone 13', 15.99, 'SKU-COQ-001', 120, 'https://images.unsplash.com/photo-1598327105666-5b89351aff97', 1, '{"couleur": "noir", "compatibilité": "iPhone 13"}', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),
(2, 'Étui Portefeuille Galaxy S22', 'Étui en cuir avec rangement cartes pour Samsung Galaxy S22', 24.90, 'SKU-COQ-002', 80, 'https://images.unsplash.com/photo-1585386959984-a4155220d6da', 1, '{"couleur": "marron", "matière": "cuir"}', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),

-- Écouteurs
(3, 'Écouteurs Bluetooth TWS', 'Écouteurs sans fil avec boîtier de charge, autonomie 20h', 39.99, 'SKU-ECO-001', 50, 'https://images.unsplash.com/photo-1583337130417-3346a1f4f1f4', 2, '{"connectivité": "Bluetooth 5.0", "autonomie": "20h"}', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),
(4, 'Écouteurs filaires Jack 3.5mm', 'Écouteurs classiques avec micro et bouton pause', 9.99, 'SKU-ECO-002', 200, 'https://images.unsplash.com/photo-1585386958707-f92efb1d6e0e', 2, '{"connecteur": "Jack 3.5mm", "longueur câble": "1.2m"}', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),

-- Chargeurs et câbles
(5, 'Chargeur USB-C 20W', 'Chargeur rapide compatible iPhone et Android', 18.50, 'SKU-CHA-001', 75, 'https://images.unsplash.com/photo-1611078489935-0cb8e440d168', 3, '{"puissance": "20W", "type": "USB-C"}', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),
(6, 'Câble Lightning MFI 1m', 'Câble certifié Apple MFI pour iPhone/iPad', 14.00, 'SKU-CHA-002', 110, 'https://images.unsplash.com/photo-1603791440384-56cd371ee9a7', 3, '{"longueur": "1m", "certification": "MFI"}', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),

-- Protections d’écran
(7, 'Verre trempé iPhone 13', 'Protection d’écran 9H ultra claire, pose facile', 7.99, 'SKU-PRO-001', 140, 'https://images.unsplash.com/photo-1626198856013-95bfe11b3943', 4, '{"dureté": "9H", "compatibilité": "iPhone 13"}', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),

-- Supports et docks
(8, 'Support Voiture Magnétique', 'Fixation grille d’aération, compatible tous smartphones', 12.49, 'SKU-SUP-001', 90, 'https://images.unsplash.com/photo-1611581874818-51c7d58230e3', 5, '{"type": "magnétique", "rotation": "360°"}', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL),
(9, 'Station de recharge 3-en-1', 'Dock pour iPhone, Apple Watch et AirPods', 59.00, 'SKU-SUP-002', 40, 'https://images.unsplash.com/photo-1616628182505-d9b443b0e3ed', 5, '{"chargeurs": "Qi", "compatibilité": "Apple"}', TRUE, '2025-04-25 12:00:00+00', '2025-04-25 12:00:00+00', NULL);

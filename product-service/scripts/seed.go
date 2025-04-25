// scripts/seed.go
package main

import (
	"fmt"
	"log"
	"os"
	"phone-accessories/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Paramètres de connexion pour la base de données dans le conteneur Docker
	// Note: Nous utilisons "localhost" car le script est exécuté depuis l'hôte
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "product_service"
	
	// Construction de la chaîne de connexion à la base de données avec paramètre client_encoding
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable client_encoding=UTF8",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	
	fmt.Println("Connexion à la base de données PostgreSQL...")
	fmt.Printf("Hôte: %s, Port: %s, Base de données: %s\n", dbHost, dbPort, dbName)
	
	// Connexion à la base de données
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible de se connecter à la base de données: %v", err)
	}
	
	fmt.Println("Connexion à la base de données réussie.")

	// Migration des tables
	fmt.Println("Migration des tables...")
	err = db.AutoMigrate(&models.Category{}, &models.Product{})
	if err != nil {
		log.Fatalf("Impossible de migrer les tables: %v", err)
	}

	// Vérification si des données existent déjà
	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		fmt.Println("La base de données contient déjà des données. Souhaitez-vous la réinitialiser ? (y/n)")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Opération annulée.")
			os.Exit(0)
		}

		// Suppression des données existantes
		fmt.Println("Suppression des données existantes...")
		// Désactiver les contraintes de clés étrangères temporairement
		db.Exec("SET CONSTRAINTS ALL DEFERRED")
		// Supprimer les données des tables
		db.Exec("DELETE FROM products")
		db.Exec("DELETE FROM categories")
		// Réinitialiser les séquences d'auto-incrémentation
		db.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
		db.Exec("ALTER SEQUENCE categories_id_seq RESTART WITH 1")
		fmt.Println("Données existantes supprimées.")
	}

	// Création des catégories
	categories := []models.Category{
		{
			Name:        "Coques et Protection",
			Description: "Protégez votre téléphone avec des coques et films protecteurs",
			ImageURL:    "https://example.com/images/categories/cases.jpg",
			IsActive:    true,
		},
		{
			Name:        "Chargeurs et Câbles",
			Description: "Câbles de charge, adaptateurs secteur et chargeurs sans fil",
			ImageURL:    "https://example.com/images/categories/chargers.jpg",
			IsActive:    true,
		},
		{
			Name:        "Audio",
			Description: "Écouteurs, casques et haut-parleurs pour votre téléphone",
			ImageURL:    "https://example.com/images/categories/audio.jpg",
			IsActive:    true,
		},
		{
			Name:        "Support et Fixation",
			Description: "Supports pour voiture, bureau et autres accessoires de fixation",
			ImageURL:    "https://example.com/images/categories/mounts.jpg",
			IsActive:    true,
		},
		{
			Name:        "Batteries externes",
			Description: "Batteries portables pour recharger vos appareils en déplacement",
			ImageURL:    "https://example.com/images/categories/powerbanks.jpg",
			IsActive:    true,
		},
	}

	// Insertion des catégories
	for i := range categories {
		result := db.Create(&categories[i])
		if result.Error != nil {
			log.Fatalf("Erreur lors de la création de la catégorie %s: %v", categories[i].Name, result.Error)
		}
		fmt.Printf("Catégorie créée: %s (ID: %d)\n", categories[i].Name, categories[i].ID)
	}

	// Création des produits
	products := []models.Product{
		{
			Name:        "Coque Silicone Premium iPhone 15",
			Description: "Coque en silicone de haute qualité pour iPhone 15, absorbe les chocs et protège votre téléphone des rayures",
			Price:       29.99,
			SKU:         "CASE-IP15-S001",
			StockLevel:  150,
			ImageURL:    "https://example.com/images/products/iphone15-case.jpg",
			CategoryID:  categories[0].ID,
			Attributes: models.JSON{
				"color":      "Noir",
				"material":   "Silicone",
				"compatible": "iPhone 15",
				"features":   []string{"Antichoc", "Antirayures", "Toucher doux"},
			},
			IsActive: true,
		},
		{
			Name:        "Film Protection Écran Samsung Galaxy S23",
			Description: "Film protecteur d'écran en verre trempé 9H pour Samsung Galaxy S23, installation facile et sans bulles",
			Price:       15.99,
			SKU:         "SCRN-SG23-G001",
			StockLevel:  200,
			ImageURL:    "https://example.com/images/products/s23-screen-protector.jpg",
			CategoryID:  categories[0].ID,
			Attributes: models.JSON{
				"material":   "Verre trempé",
				"hardness":   "9H",
				"compatible": "Samsung Galaxy S23",
				"features":   []string{"Anti-traces", "Ultra-clair", "Installation facile"},
			},
			IsActive: true,
		},
		{
			Name:        "Chargeur Rapide USB-C 30W",
			Description: "Chargeur mural USB-C avec technologie Power Delivery pour une charge rapide et efficace",
			Price:       24.99,
			SKU:         "CHRG-PD30-W001",
			StockLevel:  120,
			ImageURL:    "https://example.com/images/products/usbc-charger-30w.jpg",
			CategoryID:  categories[1].ID,
			Attributes: models.JSON{
				"power":     "30W",
				"type":      "USB-C",
				"protocol":  "Power Delivery",
				"features":  []string{"Charge rapide", "Protection contre surchauffe", "Compact"},
				"warranty":  "2 ans",
				"universal": true,
			},
			IsActive: true,
		},
		{
			Name:        "Câble USB-C vers Lightning 2m",
			Description: "Câble de charge et synchronisation tressé de 2 mètres pour appareils Apple",
			Price:       19.99,
			SKU:         "CABLE-CL2M-B001",
			StockLevel:  180,
			ImageURL:    "https://example.com/images/products/usbc-lightning-cable.jpg",
			CategoryID:  categories[1].ID,
			Attributes: models.JSON{
				"length":    "2m",
				"connector": "USB-C vers Lightning",
				"material":  "Nylon tressé",
				"color":     "Noir",
				"features":  []string{"Charge rapide", "Résistant", "MFi certifié"},
			},
			IsActive: true,
		},
		{
			Name:        "Écouteurs Bluetooth Sans Fil",
			Description: "Écouteurs intra-auriculaires sans fil avec réduction de bruit active et autonomie de 6 heures",
			Price:       89.99,
			SKU:         "AUDIO-TWS-B001",
			StockLevel:  75,
			ImageURL:    "https://example.com/images/products/wireless-earbuds.jpg",
			CategoryID:  categories[2].ID,
			Attributes: models.JSON{
				"type":        "True Wireless",
				"battery":     "6 heures",
				"waterproof":  "IPX5",
				"microphone":  true,
				"noise_canc":  true,
				"connectivity": "Bluetooth 5.2",
				"features":    []string{"Réduction de bruit", "Commandes tactiles", "Assistant vocal"},
			},
			IsActive: true,
		},
		{
			Name:        "Support Voiture Magnétique",
			Description: "Support téléphone magnétique pour tableau de bord ou grille d'aération",
			Price:       15.99,
			SKU:         "MOUNT-CAR-M001",
			StockLevel:  120,
			ImageURL:    "https://example.com/images/products/car-mount.jpg",
			CategoryID:  categories[3].ID,
			Attributes: models.JSON{
				"type":       "Magnétique",
				"mounting":   "Grille d'aération/Tableau de bord",
				"compatible": "Universel",
				"features":   []string{"Rotation 360°", "Installation facile", "Forte adhérence"},
			},
			IsActive: true,
		},
		{
			Name:        "Batterie Externe 20000mAh",
			Description: "Powerbank haute capacité avec charge rapide et deux ports USB",
			Price:       45.99,
			SKU:         "PWBNK-20K-B001",
			StockLevel:  60,
			ImageURL:    "https://example.com/images/products/powerbank.jpg",
			CategoryID:  categories[4].ID,
			Attributes: models.JSON{
				"capacity":  "20000mAh",
				"ports":     "2x USB-A, 1x USB-C",
				"input":     "USB-C 18W",
				"output":    "USB-A 5V/2.4A, USB-C 20W",
				"features":  []string{"Charge rapide", "Indicateur LED", "Charge multiple"},
				"weight":    "349g",
				"dimension": "15 x 7.5 x 2.5 cm",
			},
			IsActive: true,
		},
		{
			Name:        "Support Bureau Ajustable",
			Description: "Support de bureau réglable pour smartphones et tablettes jusqu'à 10 pouces",
			Price:       21.99,
			SKU:         "MOUNT-DESK-A001",
			StockLevel:  90,
			ImageURL:    "https://example.com/images/products/desk-stand.jpg",
			CategoryID:  categories[3].ID,
			Attributes: models.JSON{
				"material":   "Aluminium",
				"adjustable": true,
				"compatible": "Smartphones et tablettes jusqu'à 10\"",
				"angle":      "0-100°",
				"color":      "Argent",
				"features":   []string{"Pliable", "Antidérapant", "Stabilité renforcée"},
			},
			IsActive: true,
		},
		{
			Name:        "Casque Bluetooth Supra-Auriculaire",
			Description: "Casque sans fil avec réduction de bruit active et autonomie de 30 heures",
			Price:       129.99,
			SKU:         "AUDIO-HDPHN-B001",
			StockLevel:  40,
			ImageURL:    "https://example.com/images/products/bluetooth-headphones.jpg",
			CategoryID:  categories[2].ID,
			Attributes: models.JSON{
				"type":         "Supra-auriculaire",
				"connectivity": "Bluetooth 5.0",
				"battery":      "30 heures",
				"noise_canc":   true,
				"foldable":     true,
				"color":        "Noir",
				"features":     []string{"Son haute définition", "Micro intégré", "Commandes tactiles"},
			},
			IsActive: true,
		},
		{
			Name:        "Chargeur Sans Fil 15W",
			Description: "Station de charge à induction rapide compatible Qi pour smartphones",
			Price:       34.99,
			SKU:         "CHRG-WIRL-15W",
			StockLevel:  85,
			ImageURL:    "https://example.com/images/products/wireless-charger.jpg",
			CategoryID:  categories[1].ID,
			Attributes: models.JSON{
				"power":      "15W",
				"technology": "Qi",
				"design":     "Station",
				"compatible": "iPhone, Samsung, Xiaomi, etc.",
				"input":      "USB-C",
				"features":   []string{"Charge rapide", "Indicateur LED", "Protection contre surchauffe"},
			},
			IsActive: true,
		},
	}

	// Insertion des produits
	for i := range products {
		result := db.Create(&products[i])
		if result.Error != nil {
			log.Fatalf("Erreur lors de la création du produit %s: %v", products[i].Name, result.Error)
		}
		fmt.Printf("Produit créé: %s (ID: %d)\n", products[i].Name, products[i].ID)
	}

	fmt.Println("Initialisation des données terminée avec succès !")
	fmt.Printf("Total: %d catégories et %d produits créés.\n", len(categories), len(products))
}
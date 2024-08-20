package productrepository

import (
	"multilanguage/models"
	"multilanguage/repository"
)

type ProductRepository struct {
	repo repository.Repository
}

func NewProductRepository(repo repository.Repository) ProductRepository {
	return ProductRepository{
		repo: repo,
	}
}

func (r ProductRepository) FindListProduct(language string) ([]models.ProductModels, error) {
	var products []models.ProductModels

	query := `
		SELECT
			p.id,
			COALESCE(pt_name.translation, p.name) AS name,
			COALESCE(pt_desc.translation, p.description) AS description,
			p.price,
			p.quantity
		FROM product p
		LEFT JOIN translations pt_name ON pt_name.entity_type = 'product' AND pt_name.entity_id = p.id AND pt_name.language = $1 AND pt_name.field_name = 'name'
		LEFT JOIN translations pt_desc ON pt_desc.entity_type = 'product' AND pt_desc.entity_id = p.id AND pt_desc.language = $1 AND pt_desc.field_name = 'description'
    `
	rows, err := r.repo.DB.Query(query, language)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var row models.ProductModels
		err := rows.Scan(&row.ID, &row.Name, &row.Description, &row.Price, &row.Quantity)
		if err != nil {
			return products, err
		}
		products = append(products, row)
	}
	if err = rows.Err(); err != nil {
		return products, err
	}
	return products, nil
}

func (r ProductRepository) CreateProduct(req models.ProductModels) (int64, error) {
	var ID int64
	query := `
        INSERT INTO product 
            (name, description, price, quantity)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	err := r.repo.DB.QueryRow(query, req.Name, req.Description, req.Price, req.Quantity).Scan(&ID)
	if err != nil {
		return ID, err
	}
	return ID, nil
}

func (r ProductRepository) AddTranslation(req models.TranslationCreateRequest) error {
	query := `
        INSERT INTO translations 
            (entity_type, entity_id, language, field_name, translation)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (entity_type, entity_id, language, field_name)
        DO UPDATE SET translation = EXCLUDED.translation;
    `

	_, err := r.repo.DB.Exec(query, req.EntityType, req.EntityID, req.Language, req.FieldName, req.Translation)
	if err != nil {
		return err
	}
	return nil
}

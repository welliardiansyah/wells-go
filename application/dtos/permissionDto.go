package dtos

type CreatePermissionRequest struct {
	Name      string `json:"name" binding:"required"`
	CanCreate bool   `json:"can_create"`
	CanRead   bool   `json:"can_read"`
	CanUpdate bool   `json:"can_update"`
	CanDelete bool   `json:"can_delete"`
	CanExport bool   `json:"can_export"`
	CanImport bool   `json:"can_import"`
	CanView   bool   `json:"can_view"`
}

type UpdatePermissionRequest struct {
	Name      string `json:"name" binding:"required"`
	CanCreate bool   `json:"can_create"`
	CanRead   bool   `json:"can_read"`
	CanUpdate bool   `json:"can_update"`
	CanDelete bool   `json:"can_delete"`
	CanExport bool   `json:"can_export"`
	CanImport bool   `json:"can_import"`
	CanView   bool   `json:"can_view"`
}

type PermissionResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CanCreate bool   `json:"can_create"`
	CanRead   bool   `json:"can_read"`
	CanUpdate bool   `json:"can_update"`
	CanDelete bool   `json:"can_delete"`
	CanExport bool   `json:"can_export"`
	CanImport bool   `json:"can_import"`
	CanView   bool   `json:"can_view"`
}

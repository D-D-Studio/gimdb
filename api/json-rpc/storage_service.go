package json_rpc

import (
	"gimdb/domain"
	"log"
	"net/http"
)

type StorageService struct {
	workspace *domain.Workspace
}

func NewStorageService(workspace *domain.Workspace) *StorageService {
	storageService := new(StorageService)

	storageService.workspace = workspace

	return storageService
}

type CreateStorageArgs struct {
	Name string `json:"name"`
}

func (s *StorageService) CreateStorage(r *http.Request, args *CreateStorageArgs, reply *any) error {
	log.Printf("[JSON-RPC] CREATE_STORAGE \"%s\"\n", args.Name)

	return s.workspace.CreateStorage(args.Name)
}

type DeleteStorageArgs struct {
	Name string `json:"name"`
}

func (s *StorageService) DeleteStorage(r *http.Request, args *DeleteStorageArgs, reply *any) error {
	log.Printf("[JSON-RPC] DELETE_STORAGE \"%s\"\n", args.Name)

	return s.workspace.DeleteStorage(args.Name)
}

type ExistsStorageArgs struct {
	Name string `json:"name"`
}

type ExistsStorageReply struct {
	IsExists bool `json:"isExists"`
}

func (s *StorageService) ExistsStorage(r *http.Request, args *ExistsStorageArgs, reply *ExistsStorageReply) error {
	log.Printf("[JSON-RPC] EXISTS_STORAGE \"%s\"\n", args.Name)

	reply.IsExists = s.workspace.IsStorageExists(args.Name)

	return nil
}

type GetArgs struct {
	Storage string `json:"storage"`
	Key     string `json:"key"`
}

type GetReply struct {
	Value string `json:"value"`
}

func (s *StorageService) Get(r *http.Request, args *GetArgs, reply *GetReply) error {
	log.Printf("[JSON-RPC] GET \"%s\" \"%s\"\n", args.Storage, args.Key)

	storage, err := s.workspace.GetStorage(args.Storage)

	if err != nil {
		return err
	}

	value, err := storage.Get(args.Key)

	if err != nil {
		return err
	}

	reply.Value = value

	return nil
}

type SetArgs struct {
	Storage string `json:"storage"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}

func (s *StorageService) Set(r *http.Request, args *SetArgs, reply *any) error {
	log.Printf("[JSON-RPC] SET \"%s\" \"%s\" \"%s\"\n", args.Storage, args.Key, args.Value)

	storage, err := s.workspace.GetStorage(args.Storage)

	if err != nil {
		return err
	}

	storage.Set(args.Key, args.Value)

	return nil
}

type DeleteArgs struct {
	Storage string `json:"storage"`
	Key     string `json:"key"`
}

func (s *StorageService) Delete(r *http.Request, args *DeleteArgs, reply *any) error {
	log.Printf("[JSON-RPC] DELETE \"%s\" \"%s\"\n", args.Storage, args.Key)

	storage, err := s.workspace.GetStorage(args.Storage)

	if err != nil {
		return err
	}

	return storage.Delete(args.Key)
}

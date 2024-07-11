package mappers

import (
	"health-check/config"
	"health-check/domain"
	"health-check/errors"
	"log"

	magnetarapi "github.com/c12s/magnetar/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapError(err *errors.ErrorStruct) error {
	if err == nil {
		log.Println("Received nil error in MapError")
		return nil
	}
	log.Println("Error from gRPC mapper:", err)
	return status.Error(codes.Code(err.GetErrorStatus()), err.GetErrorMessage())
}

func MapFromApiExternalApplicationToModelExternalApplication(list *magnetarapi.ListAllNodesResp, nodsConfig *config.NodeConfig) {
	//if len(nodsConfig.GetLoadedIDs()) == 0 {
	//	for _, node := range list.GetNodes() {
	//		nodsConfig.AppendLoadedIDs(node.GetId())
	//		nodsConfig.AppendNewNode(node.GetId())
	//	}
	//	return
	//}
	//listOfNewNodes := make([]string, 0)
	//for _, node := range list.GetNodes() {
	//	listOfNewNodes = append(listOfNewNodes, node.GetId())
	//}
	//newListMap := utils.ConvertFromStringArrayToMap(listOfNewNodes)
	//currentListMap := utils.ConvertFromStringArrayToMap(nodsConfig.GetLoadedIDs())
	//for id := range currentListMap {
	//	if !newListMap[id] {
	//		delete(currentListMap, id)
	//		nodsConfig.RemoveNode(id)
	//	}
	//}
	//
	//for id := range newListMap {
	//	if !currentListMap[id] {
	//		nodsConfig.AppendNewNode(id)
	//	}
	//}
	newMapWithNodes := make(map[string]domain.Node)
	for _, node := range list.GetNodes() {
		newMapWithNodes[node.GetId()] = *domain.NewNode(node.GetId())

	}
	nodsConfig.SetNodes(newMapWithNodes)
}

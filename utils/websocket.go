package utils

// import (
// 	"github.com/gorilla/websocket"
// )

// func WSServiceHealthCheck(conn *websocket.Conn) error {
// 	// send health check message to accounts service
// 	err := conn.WriteMessage(websocket.TextMessage, []byte(`{"action": "healthCheck", "data": {}}`))
// 	if err != nil {
// 		log.Println("Failed to send health check message to Accounts Service: ", err)
// 		ProcessLogMessage(Error, "Failed to send health check message to Accounts Service: "+err.Error())
// 		return err
// 	}
// 	// Read the response from accounts service
// 	_, response, _ := conn.ReadMessage()
// 	ProcessLogMessage(Info, "Accounts Service HealthCheck: "+string(response))

// 	// Unmarshal response
// 	var resBody map[string]interface{}
// 	err = json.Unmarshal(response, &resBody)
// 	if err != nil {
// 		log.Println("Failed to unmarshal response from Accounts Service: ", err)
// 		ProcessLogMessage(Error, "Failed to unmarshal response from Accounts Service: "+err.Error())
// 		return err
// 	}

// 	// Check for errors in response
// 	if resBody["error"] != nil {
// 		log.Println("Error in response from Accounts Service: ", resBody["error"])
// 		ProcessLogMessage(Error, "Error in response from Accounts Service: "+resBody["error"].(string))
// 		return err
// 	}

// 	return nil
// }

// func ConnectService(url string, serviceName string) (*websocket.Conn, error) {
// 	// Connect to Accounts Service
// 	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
// 	if err != nil {
// 		ProcessLogMessage(Error, "Failed to connect to "+serviceName+" Service: "+err.Error())
// 		log.Println("Failed to connect to "+serviceName+" Service: ", err)
// 		return nil, err
// 	}
// 	// defer conn.Close()

// 	// perform a ping to the service
// 	err = conn.WriteMessage(websocket.PingMessage, nil)
// 	if err != nil {
// 		ProcessLogMessage(Error, "Failed to ping service: "+err.Error())
// 		log.Println("Failed to ping service: ", err)
// 		return nil, err
// 	}

// 	// health check
// 	health := WSServiceHealthCheck(conn)
// 	if health != nil {
// 		ProcessLogMessage(Error, "Failed HealthCheck to "+serviceName+" Service: "+health.Error())
// 		log.Println("Failed HealthCheck to "+serviceName+" Service: ", health.Error())
// 		return nil, health
// 	}
// 	return conn, nil
// }

// func ReConnectService(SrvConfig *ServConfig) (*websocket.Conn, error) {
// 	// Connect to Service
// 	conn, _, err := websocket.DefaultDialer.Dial(SrvConfig.URL, nil)
// 	if err != nil {
// 		ProcessLogMessage(Error, "Failed to connect to "+SrvConfig.SrvName+" Service: "+err.Error())
// 		log.Println("Failed to connect to "+SrvConfig.SrvName+" Service: ", err.Error())
// 		return nil, err
// 	}
// 	// defer conn.Close()

// 	// perform a ping to the service
// 	err = conn.WriteMessage(websocket.PingMessage, nil)
// 	if err != nil {
// 		ProcessLogMessage(Error, "Failed to ping service: "+err.Error())
// 		log.Println("Failed to ping service: ", err)
// 		return nil, err
// 	}

// 	// health check
// 	health := WSServiceHealthCheck(conn)
// 	if health != nil {
// 		ProcessLogMessage(Error, "Failed HealthCheck to "+SrvConfig.SrvName+" Service: "+health.Error())
// 		log.Println("Failed HealthCheck to "+SrvConfig.SrvName+" Service: ", health.Error())
// 		return nil, health
// 	}

// 	return conn, nil
// }

// func MakeRequest(serviceConfig *ServConfig, reqData Request) (map[string]interface{}, error) {
// 	jsonData, err := json.Marshal(reqData)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = serviceConfig.Conn.WriteMessage(websocket.TextMessage, jsonData)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Read the response
// 	_, response, err := serviceConfig.Conn.ReadMessage()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Unmarshal response
// 	var resBody map[string]interface{}
// 	err = json.Unmarshal(response, &resBody)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check for errors in response
// 	if resBody["error"] != nil {
// 		return nil, errors.New(resBody["error"].(string))
// 	}

// 	return resBody, nil
// }

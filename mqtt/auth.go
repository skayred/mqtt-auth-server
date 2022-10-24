package mqtt

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var ControlTopic = "$CONTROL/dynamic-security/v1"

func getAdminClient(brokerIP string, brokerPort string, username string, password string) (mqtt.Client, error) {
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", brokerIP, port))
	opts.SetClientID("mqtt-auth-server")
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())

		return nil, token.Error()
	}

	return client, nil
}

// This methods creates the role under the name of user login, and then creates an account under that role. Token is the MQTT password
func CreateUser(brokerIP string, brokerPort string, username string, password string, login string, token string) error {
	var client, err = getAdminClient(brokerIP, brokerPort, username, password)

	if err != nil {
		return err
	}

	defer client.Disconnect(250)

	bytes, err := json.Marshal(EnvelopeRole{
		Commands: []*CreateRoleReq{{
			Command:  "createRole",
			Rolename: &login,
			ACLs:     []*ACL{},
		},
		}})

	if err != nil {
		fmt.Println(err)

		return err
	}

	cancellationToken := client.Publish(ControlTopic, 0, false, bytes)
	cancellationToken.Wait()

	bytes, err = json.Marshal(
		EnvelopeClient{
			Commands: []*CreateClientReq{
				{
					Command:  "createClient",
					Username: &login,
					Password: &token,
					Roles: []*Role{
						{
							Rolename: &login,
						},
					},
				}},
		})

	if err != nil {
		fmt.Println(err)

		return err
	}

	cancellationToken = client.Publish(ControlTopic, 0, false, bytes)
	cancellationToken.Wait()

	return nil
}

// This method allows to update the user MQQT account name and token (if it needs to be regenerated)
func UpdateUser(brokerIP string, brokerPort string, username string, password string, login string, token string) error {
	var client, err = getAdminClient(brokerIP, brokerPort, username, password)

	if err != nil {
		return err
	}

	defer client.Disconnect(250)

	bytes, err := json.Marshal(
		EnvelopeClient{
			Commands: []*CreateClientReq{
				{
					Command:  "modifyClient",
					Username: &login,
					Password: &token,
					Roles: []*Role{
						{
							Rolename: &login,
						},
					},
				}},
		})

	if err != nil {
		fmt.Println(err)

		return err
	}

	cancellationToken := client.Publish(ControlTopic, 0, false, bytes)
	cancellationToken.Wait()

	return nil
}

// This methods grants the user role and, therefore, user account a subscribe permisions for the listed topics
func GrantTopicAccess(brokerIP string, brokerPort string, username string, password string, login string, topics []string) error {
	var client, err = getAdminClient(brokerIP, brokerPort, username, password)

	if err != nil {
		return err
	}

	defer client.Disconnect(250)

	var topicsList = []*ChangeACLReq{}

	for _, topicName := range topics {
		topicsList = append(topicsList, &ChangeACLReq{
			Command:  "addRoleACL",
			Rolename: &login,
			ACLtype:  "subscribeLiteral",
			Topic:    topicName,
			Priority: 1,
			Allow:    true,
		})
	}

	bytes, err := json.Marshal(EnvelopeACL{
		Commands: topicsList,
	})

	if err != nil {
		fmt.Println(err)

		return err
	} else {
		fmt.Println("publishing", string(bytes))
	}

	cancellationToken := client.Publish(ControlTopic, 0, false, bytes)
	cancellationToken.Wait()

	return nil
}

// This methods removes the access ro topics subscriptions from the provided user
func DenyTopicAccess(brokerIP string, brokerPort string, username string, password string, login string, topics []string) error {
	var client, err = getAdminClient(brokerIP, brokerPort, username, password)

	if err != nil {
		return err
	}

	defer client.Disconnect(250)

	var topicsList = []*ChangeACLReq{}

	for _, topicName := range topics {
		topicsList = append(topicsList, &ChangeACLReq{
			Command:  "removeRoleACL",
			Rolename: &login,
			ACLtype:  "subscribeLiteral",
			Topic:    topicName,
		})
	}

	bytes, err := json.Marshal(EnvelopeACL{
		Commands: topicsList,
	})

	if err != nil {
		fmt.Println(err)

		return err
	} else {
		fmt.Println("publishing", string(bytes))
	}

	cancellationToken := client.Publish(ControlTopic, 0, false, bytes)
	cancellationToken.Wait()

	return nil
}

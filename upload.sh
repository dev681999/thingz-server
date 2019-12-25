sudo scp -i server_key.pem build/api admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/
# sudo scp -i server_key.pem build/config.json admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/
# sudo scp -i server_key.pem build/docker-compose.yaml admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/
sudo scp -i server_key.pem build/mqtt admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/
# sudo scp -i server_key.pem build/nats-server admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/
sudo scp -i server_key.pem build/project admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/
sudo scp -i server_key.pem build/rule admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/
# sudo scp -i server_key.pem build/run.sh admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/
sudo scp -i server_key.pem build/thing admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/
sudo scp -i server_key.pem build/user admin@ec2-54-191-142-124.us-west-2.compute.amazonaws.com:/var/thingz/

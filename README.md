# load-balancing-algos

Load-balancing algos implementation written in Golang.

## Algorithms presented:

* Round-robin
* Weighted round-robin (TBD)
* Least Connection (TBD)
* Fixed Weighting (TBD)
* Weighted Response Time (TBD)
* Source IP Hash (TBD)
* URL Hash (TBD)

## Components

* [Echo-server](echo-server/README.md)
* [Load-balancer](load-balancer/README.md)

## Installation and running

* Setup Docker images in your local hub - run `./build-images.sh`
* Run 3 echo-servers and 1 lb: `docker-compose up`
* Switch over algorithms used via `curl -X POST http://127.0.0.1:8080/lb/switch?type=TYPE` where `TYPE` is presented in
  the table below:

| Name  | Parameter |
| ------------- | ------------- |
| Round-robin  | round_robin  |
| Weighted round-robin  | weighted_round_robin  |
| Least Connection  | least_connection  |
| Fixed Weighting  | fixed_weighting  |
| Weighted Response Time  | weighted_response  |
| Source IP Hash  | ip_hash  |
| URL Hash  | url_hash  |

* Test via executing `curl -X GET http://127.0.0.1:8080/hello`

## Authors

**Vladislav Sidlyarevich** - [Github profile](https://github.com/vlsidlyarevich)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
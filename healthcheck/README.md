# Health Check API

You have applied the Microservice architecture pattern. Sometimes a service instance can be incapable of handling requests yet still be running. For example, it might have ran out of database connections. When this occurs, the monitoring system should generate a alert. Also, the load balancer or service registry should not route requests to the failed service instance.

## Problem

How to detect that a running service instance is unable to handle requests?

## Forces

- An alert should be generated when a service instance fails
- Requests should be routed to working service instances

## Solution

A service has an health check API endpoint (e.g HTTP /health) that returns the health of the service. The API endpoint handler performs various checks, such as 

- The status of the connections to the infrastructure services used by the service instance
- The status of the host 
- Application specific logic

## Usage

Base on: https://github.com/hellofresh/health-go

Implements some generic checkers for the following services
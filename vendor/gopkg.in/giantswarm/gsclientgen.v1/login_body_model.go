/*
 * Giant Swarm legacy API
 *
 * Caution: This is an incomplete description of some legacy API functions.
 *
 * OpenAPI spec version: legacy
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package gsclientgen

type LoginBodyModel struct {

	// base64 encoded password
	Password string `json:"password"`
}
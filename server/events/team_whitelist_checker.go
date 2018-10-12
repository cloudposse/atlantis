// Copyright 2017 HootSuite Media Inc.
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Modified hereafter by contributors to runatlantis/atlantis.

package events

import (
	"strings"
)

// Wildcard matches all teams and all commands
const wildcard = "*"

// mapOfStrings is an alias for map[string]string
type mapOfStrings map[string]string

// TeamWhitelistChecker implements checking the teams and the operations that the members
// of a particular team are allowed to perform
type TeamWhitelistChecker struct {
	rules []mapOfStrings
}

// NewTeamWhitelistChecker constructs a new checker
func NewTeamWhitelistChecker(whitelist string) (*TeamWhitelistChecker, error) {
	var rules []mapOfStrings
	pairs := strings.Split(whitelist, ",")
	for _, pair := range pairs {
		values := strings.Split(pair, ":")
		team := strings.TrimSpace(values[0])
		command := strings.TrimSpace(values[1])
		m := mapOfStrings{team: command}
		rules = append(rules, m)
	}
	return &TeamWhitelistChecker{
		rules: rules,
	}, nil
}

// IsCommandAllowedForTeam returns true if the team is allowed to execute the command
// and false otherwise.
func (checker *TeamWhitelistChecker) IsCommandAllowedForTeam(team string, command string) bool {
	t := strings.TrimSpace(team)
	c := strings.TrimSpace(command)
	for _, rule := range checker.rules {
		for key, value := range rule {
			if (key == wildcard || strings.EqualFold(key, t)) && (value == wildcard || strings.EqualFold(value, c)) {
				return true
			}
		}
	}
	return false
}

// IsCommandAllowedForAnyTeam returns true if any of the teams is allowed to execute the command
// and false otherwise.
func (checker *TeamWhitelistChecker) IsCommandAllowedForAnyTeam(teams []string, command string) bool {
	c := strings.TrimSpace(command)
	if teams == nil || len(teams) == 0 {
		for _, rule := range checker.rules {
			for key, value := range rule {
				if (key == wildcard) && (value == wildcard || strings.EqualFold(value, c)) {
					return true
				}
			}
		}
	} else {
		for _, t := range teams {
			if checker.IsCommandAllowedForTeam(t, command) {
				return true
			}
		}
	}
	return false
}

/*
Copyright (C) 2022 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package state

import (
	"strings"

	"k8s.io/apimachinery/pkg/labels"
)

func (f *Fetcher) getAccessControlPolicies() (map[string]*AccessControlPolicy, error) {
	policies, err := f.hub.Hub().V1alpha1().AccessControlPolicies().Lister().List(labels.Everything())
	if err != nil {
		return nil, err
	}

	result := make(map[string]*AccessControlPolicy)
	for _, policy := range policies {
		acp := &AccessControlPolicy{
			Name:      policy.Name,
			Namespace: policy.Namespace,
		}

		switch {
		case policy.Spec.JWT != nil:
			acp.Method = "jwt"
			acp.JWT = &AccessControlPolicyJWT{
				SigningSecretBase64Encoded: policy.Spec.JWT.SigningSecretBase64Encoded,
				PublicKey:                  policy.Spec.JWT.PublicKey,
				JWKsFile:                   policy.Spec.JWT.JWKsFile,
				JWKsURL:                    policy.Spec.JWT.JWKsURL,
				StripAuthorizationHeader:   policy.Spec.JWT.StripAuthorizationHeader,
				ForwardHeaders:             policy.Spec.JWT.ForwardHeaders,
				TokenQueryKey:              policy.Spec.JWT.TokenQueryKey,
				Claims:                     policy.Spec.JWT.Claims,
			}

			if policy.Spec.JWT.SigningSecret != "" {
				acp.JWT.SigningSecret = "redacted"
			}
		case policy.Spec.BasicAuth != nil:
			acp.Method = "basicauth"
			acp.BasicAuth = &AccessControlPolicyBasicAuth{
				Users:                    redactPasswords(policy.Spec.BasicAuth.Users),
				Realm:                    policy.Spec.BasicAuth.Realm,
				StripAuthorizationHeader: policy.Spec.BasicAuth.StripAuthorizationHeader,
				ForwardUsernameHeader:    policy.Spec.BasicAuth.ForwardUsernameHeader,
			}
		case policy.Spec.OIDC != nil:
			acp.Method = "oidc"
			acp.OIDC = &AccessControlPolicyOIDC{
				Issuer:   policy.Spec.OIDC.Issuer,
				ClientID: policy.Spec.OIDC.ClientID,
				Secret: &SecretReference{
					Name:      policy.Spec.OIDC.Secret.Name,
					Namespace: policy.Spec.OIDC.Secret.Namespace,
				},
				RedirectURL: policy.Spec.OIDC.RedirectURL,
				LogoutURL:   policy.Spec.OIDC.LogoutURL,
				Scopes:      policy.Spec.OIDC.Scopes,
				AuthParams:  policy.Spec.OIDC.AuthParams,
				StateCookie: &AuthStateCookie{
					Path:     policy.Spec.OIDC.StateCookie.Path,
					Domain:   policy.Spec.OIDC.StateCookie.Domain,
					SameSite: policy.Spec.OIDC.StateCookie.SameSite,
					Secure:   policy.Spec.OIDC.StateCookie.Secure,
				},
				Session: &AuthSession{
					Path:     policy.Spec.OIDC.Session.Path,
					Domain:   policy.Spec.OIDC.Session.Domain,
					SameSite: policy.Spec.OIDC.Session.SameSite,
					Secure:   policy.Spec.OIDC.Session.Secure,
					Refresh:  policy.Spec.OIDC.Session.Refresh,
				},
				ForwardHeaders: policy.Spec.OIDC.ForwardHeaders,
				Claims:         policy.Spec.OIDC.Claims,
			}
		case policy.Spec.OIDCGoogle != nil:
			acp.Method = "oidcGoogle"
			acp.OIDCGoogle = &AccessControlPolicyOIDCGoogle{
				ClientID: policy.Spec.OIDCGoogle.ClientID,
				Secret: &SecretReference{
					Name:      policy.Spec.OIDCGoogle.Secret.Name,
					Namespace: policy.Spec.OIDCGoogle.Secret.Namespace,
				},
				RedirectURL: policy.Spec.OIDCGoogle.RedirectURL,
				LogoutURL:   policy.Spec.OIDCGoogle.LogoutURL,
				AuthParams:  policy.Spec.OIDCGoogle.AuthParams,
				StateCookie: &AuthStateCookie{
					Path:     policy.Spec.OIDCGoogle.StateCookie.Path,
					Domain:   policy.Spec.OIDCGoogle.StateCookie.Domain,
					SameSite: policy.Spec.OIDCGoogle.StateCookie.SameSite,
					Secure:   policy.Spec.OIDCGoogle.StateCookie.Secure,
				},
				Session: &AuthSession{
					Path:     policy.Spec.OIDCGoogle.Session.Path,
					Domain:   policy.Spec.OIDCGoogle.Session.Domain,
					SameSite: policy.Spec.OIDCGoogle.Session.SameSite,
					Secure:   policy.Spec.OIDCGoogle.Session.Secure,
					Refresh:  policy.Spec.OIDCGoogle.Session.Refresh,
				},
				ForwardHeaders: policy.Spec.OIDCGoogle.ForwardHeaders,
				Emails:         policy.Spec.OIDCGoogle.Emails,
			}
		default:
			continue
		}

		result[policy.Name] = acp
	}

	return result, nil
}

func redactPasswords(rawUsers []string) string {
	var users []string

	for _, u := range rawUsers {
		i := strings.Index(u, ":")
		if i <= 0 {
			continue
		}

		users = append(users, u[:i]+":redacted")
	}

	return strings.Join(users, ",")
}

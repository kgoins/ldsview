LDIF ?= dump.ldif
SEARCH = ldsview -f "${LDIF}" search
UAC = ldsview -f "${LDIF}" uac --search

query.users = objectClass:=user,objectClass:!=computer
query.computers = objectClass:=computer
query.groups = objectClass:=group
query.domain_admin = cn:~Domain Admins
query.poss_svc_accnts = sAMAccountName:~svc

query.pass_not_reqd = 32
query.pass_cant_change = 64
query.users_dont_expire = 65536
query.trusted_4_delegation = 524288
query.preauth_not_reqd = 4194304
query.password_expired = 8388608
query.trust2auth4delegation = 16777216

SEARCH_TARGETS = users \
          computers \
          groups \
          domain_admin \
          poss_svc_accnts

UAC_TARGETS = pass_not_reqd \
          pass_cant_change \
          users_dont_expire \
          trusted_4_delegation \
          preauth_not_reqd \
          password_expired \
          trust2auth4delegation

target=$(word 1, $@)

all: ${SEARCH_TARGETS} ${UAC_TARGETS}

${SEARCH_TARGETS}:
        ${SEARCH} "${query.${target}}" > "${target}.ldif"

${UAC_TARGETS}:
        ${UAC} "${query.${target}}" > "${target}.ldif"

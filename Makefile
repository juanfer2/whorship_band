# Configuration
copy-config:
	cp .env.example .env

	cp src/shared/infrastructure/persistence/postgres/config/database_env.yml src/shared/infrastructure/persistence/postgres/config/database_development.yml
	cp src/shared/infrastructure/persistence/postgres/config/database_env.yml src/shared/infrastructure/persistence/postgres/config/database_test.yml
	cp src/shared/infrastructure/persistence/postgres/config/database_env.yml src/shared/infrastructure/persistence/postgres/config/database_staging.yml

	cp src/shared/infrastructure/clients/b2chat/config/b2chat_env.example.yml src/shared/infrastructure/clients/b2chat/config/b2chat_development.yml
	cp src/shared/infrastructure/clients/b2chat/config/b2chat_env.example.yml src/shared/infrastructure/clients/b2chat/config/b2chat_test.yml
	cp src/shared/infrastructure/clients/b2chat/config/b2chat_env.example.yml src/shared/infrastructure/clients/b2chat/config/b2chat_staging.yml

	cp src/shared/infrastructure/clients/wompi/config/wompi_credentias_env.yml src/shared/infrastructure/clients/wompi/config/wompi_credentias_development.yml
	cp src/shared/infrastructure/clients/wompi/config/wompi_credentias_env.yml src/shared/infrastructure/clients/wompi/config/wompi_credentias_test.yml
	cp src/shared/infrastructure/clients/wompi/config/wompi_credentias_env.yml src/shared/infrastructure/clients/wompi/config/wompi_credentias_staging.yml

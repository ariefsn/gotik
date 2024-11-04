import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  schema: 'http://localhost:8080/query',
  documents: ['src/graphql/**/*.gql.{ts,tsx}'],
  generates: {
    'src/graphql/generated.ts': {
      plugins: ['typescript', 'typescript-operations', 'typescript-validation-schema'],
      config: {
        schema: 'zod'
      }
    }
  }
};

export default config;

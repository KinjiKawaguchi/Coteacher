import { createGrpcWebTransport } from '@bufbuild/connect-web';

export const backendGrpcBaseUrl = 'http://localhost:50051';
export const backendGrpcTransport = createGrpcWebTransport({
  baseUrl: backendGrpcBaseUrl,
});

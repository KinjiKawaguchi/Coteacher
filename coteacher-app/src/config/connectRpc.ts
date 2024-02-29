import { createGrpcWebTransport } from '@bufbuild/connect-web';

export const backendGrpcBaseUrl =
  process.env.NEXT_PUBLIC_API_ENDPOINT || 'http://localhost:8080';
export const backendGrpcTransport = createGrpcWebTransport({
  baseUrl: backendGrpcBaseUrl,
});

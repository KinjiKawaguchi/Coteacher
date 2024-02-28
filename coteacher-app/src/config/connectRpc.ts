import { createGrpcWebTransport } from '@bufbuild/connect-web';

export const backendGrpcBaseUrl =
  process.env.NEXT_PUBLIC_API_ENDPOINT || 'http://localhost:50051';
export const backendGrpcTransport = createGrpcWebTransport({
  baseUrl: backendGrpcBaseUrl,
});

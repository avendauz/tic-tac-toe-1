import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../tictactoe1/params";
export declare const protobufPackage = "avendauz.tictactoe1.tictactoe1";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryAllOpenGamesRequest {
}
export interface QueryAllOpenGamesResponse {
    games: string[];
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryAllOpenGamesRequest: {
    encode(_: QueryAllOpenGamesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllOpenGamesRequest;
    fromJSON(_: any): QueryAllOpenGamesRequest;
    toJSON(_: QueryAllOpenGamesRequest): unknown;
    fromPartial(_: DeepPartial<QueryAllOpenGamesRequest>): QueryAllOpenGamesRequest;
};
export declare const QueryAllOpenGamesResponse: {
    encode(message: QueryAllOpenGamesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllOpenGamesResponse;
    fromJSON(object: any): QueryAllOpenGamesResponse;
    toJSON(message: QueryAllOpenGamesResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllOpenGamesResponse>): QueryAllOpenGamesResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    AllOpenGames(request: QueryAllOpenGamesRequest): Promise<QueryAllOpenGamesResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    AllOpenGames(request: QueryAllOpenGamesRequest): Promise<QueryAllOpenGamesResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

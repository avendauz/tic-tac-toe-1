import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "avendauz.tictactoe1.tictactoe1";
export interface MsgOpenGame {
    creator: string;
    uuid: string;
}
export interface MsgOpenGameResponse {
}
export interface MsgAcceptGame {
    creator: string;
    uuid: string;
    challenger: string;
}
export interface MsgAcceptGameResponse {
}
export declare const MsgOpenGame: {
    encode(message: MsgOpenGame, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgOpenGame;
    fromJSON(object: any): MsgOpenGame;
    toJSON(message: MsgOpenGame): unknown;
    fromPartial(object: DeepPartial<MsgOpenGame>): MsgOpenGame;
};
export declare const MsgOpenGameResponse: {
    encode(_: MsgOpenGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgOpenGameResponse;
    fromJSON(_: any): MsgOpenGameResponse;
    toJSON(_: MsgOpenGameResponse): unknown;
    fromPartial(_: DeepPartial<MsgOpenGameResponse>): MsgOpenGameResponse;
};
export declare const MsgAcceptGame: {
    encode(message: MsgAcceptGame, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAcceptGame;
    fromJSON(object: any): MsgAcceptGame;
    toJSON(message: MsgAcceptGame): unknown;
    fromPartial(object: DeepPartial<MsgAcceptGame>): MsgAcceptGame;
};
export declare const MsgAcceptGameResponse: {
    encode(_: MsgAcceptGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAcceptGameResponse;
    fromJSON(_: any): MsgAcceptGameResponse;
    toJSON(_: MsgAcceptGameResponse): unknown;
    fromPartial(_: DeepPartial<MsgAcceptGameResponse>): MsgAcceptGameResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    OpenGame(request: MsgOpenGame): Promise<MsgOpenGameResponse>;
    AcceptGame(request: MsgAcceptGame): Promise<MsgAcceptGameResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    OpenGame(request: MsgOpenGame): Promise<MsgOpenGameResponse>;
    AcceptGame(request: MsgAcceptGame): Promise<MsgAcceptGameResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

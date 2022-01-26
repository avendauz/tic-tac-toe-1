import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "avendauz.tictactoe1.tictactoe1";
export interface MsgOpenGame {
    creator: string;
    uuid: string;
}
export interface MsgOpenGameResponse {
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
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    OpenGame(request: MsgOpenGame): Promise<MsgOpenGameResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    OpenGame(request: MsgOpenGame): Promise<MsgOpenGameResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

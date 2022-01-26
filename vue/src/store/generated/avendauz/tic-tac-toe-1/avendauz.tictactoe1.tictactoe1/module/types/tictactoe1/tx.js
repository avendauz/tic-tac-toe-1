/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "avendauz.tictactoe1.tictactoe1";
const baseMsgOpenGame = { creator: "", uuid: "" };
export const MsgOpenGame = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.uuid !== "") {
            writer.uint32(18).string(message.uuid);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgOpenGame };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.uuid = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgOpenGame };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.uuid !== undefined && object.uuid !== null) {
            message.uuid = String(object.uuid);
        }
        else {
            message.uuid = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.uuid !== undefined && (obj.uuid = message.uuid);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgOpenGame };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.uuid !== undefined && object.uuid !== null) {
            message.uuid = object.uuid;
        }
        else {
            message.uuid = "";
        }
        return message;
    },
};
const baseMsgOpenGameResponse = {};
export const MsgOpenGameResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgOpenGameResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgOpenGameResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgOpenGameResponse };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    OpenGame(request) {
        const data = MsgOpenGame.encode(request).finish();
        const promise = this.rpc.request("avendauz.tictactoe1.tictactoe1.Msg", "OpenGame", data);
        return promise.then((data) => MsgOpenGameResponse.decode(new Reader(data)));
    }
}

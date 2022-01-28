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
const baseMsgAcceptGame = { creator: "", uuid: "", challenger: "" };
export const MsgAcceptGame = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.uuid !== "") {
            writer.uint32(18).string(message.uuid);
        }
        if (message.challenger !== "") {
            writer.uint32(26).string(message.challenger);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAcceptGame };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.uuid = reader.string();
                    break;
                case 3:
                    message.challenger = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAcceptGame };
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
        if (object.challenger !== undefined && object.challenger !== null) {
            message.challenger = String(object.challenger);
        }
        else {
            message.challenger = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.uuid !== undefined && (obj.uuid = message.uuid);
        message.challenger !== undefined && (obj.challenger = message.challenger);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAcceptGame };
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
        if (object.challenger !== undefined && object.challenger !== null) {
            message.challenger = object.challenger;
        }
        else {
            message.challenger = "";
        }
        return message;
    },
};
const baseMsgAcceptGameResponse = {};
export const MsgAcceptGameResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAcceptGameResponse };
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
        const message = { ...baseMsgAcceptGameResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAcceptGameResponse };
        return message;
    },
};
const baseMsgMove = { creator: "", uuid: "", player: "", cell: 0 };
export const MsgMove = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.uuid !== "") {
            writer.uint32(18).string(message.uuid);
        }
        if (message.player !== "") {
            writer.uint32(26).string(message.player);
        }
        if (message.cell !== 0) {
            writer.uint32(32).int32(message.cell);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgMove };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.uuid = reader.string();
                    break;
                case 3:
                    message.player = reader.string();
                    break;
                case 4:
                    message.cell = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgMove };
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
        if (object.player !== undefined && object.player !== null) {
            message.player = String(object.player);
        }
        else {
            message.player = "";
        }
        if (object.cell !== undefined && object.cell !== null) {
            message.cell = Number(object.cell);
        }
        else {
            message.cell = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.uuid !== undefined && (obj.uuid = message.uuid);
        message.player !== undefined && (obj.player = message.player);
        message.cell !== undefined && (obj.cell = message.cell);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgMove };
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
        if (object.player !== undefined && object.player !== null) {
            message.player = object.player;
        }
        else {
            message.player = "";
        }
        if (object.cell !== undefined && object.cell !== null) {
            message.cell = object.cell;
        }
        else {
            message.cell = 0;
        }
        return message;
    },
};
const baseMsgMoveResponse = {};
export const MsgMoveResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgMoveResponse };
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
        const message = { ...baseMsgMoveResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgMoveResponse };
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
    AcceptGame(request) {
        const data = MsgAcceptGame.encode(request).finish();
        const promise = this.rpc.request("avendauz.tictactoe1.tictactoe1.Msg", "AcceptGame", data);
        return promise.then((data) => MsgAcceptGameResponse.decode(new Reader(data)));
    }
    Move(request) {
        const data = MsgMove.encode(request).finish();
        const promise = this.rpc.request("avendauz.tictactoe1.tictactoe1.Msg", "Move", data);
        return promise.then((data) => MsgMoveResponse.decode(new Reader(data)));
    }
}

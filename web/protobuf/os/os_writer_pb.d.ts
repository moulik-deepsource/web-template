import * as jspb from 'google-protobuf'

import * as protobuf_os_os_pb from '../../protobuf/os/os_pb';


export class CreateOneRequest extends jspb.Message {
  getRequest(): protobuf_os_os_pb.OS | undefined;
  setRequest(value?: protobuf_os_os_pb.OS): CreateOneRequest;
  hasRequest(): boolean;
  clearRequest(): CreateOneRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateOneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateOneRequest): CreateOneRequest.AsObject;
  static serializeBinaryToWriter(message: CreateOneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateOneRequest;
  static deserializeBinaryFromReader(message: CreateOneRequest, reader: jspb.BinaryReader): CreateOneRequest;
}

export namespace CreateOneRequest {
  export type AsObject = {
    request?: protobuf_os_os_pb.OS.AsObject,
  }
}

export class CreateOneResponse extends jspb.Message {
  getResult(): protobuf_os_os_pb.OS | undefined;
  setResult(value?: protobuf_os_os_pb.OS): CreateOneResponse;
  hasResult(): boolean;
  clearResult(): CreateOneResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateOneResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateOneResponse): CreateOneResponse.AsObject;
  static serializeBinaryToWriter(message: CreateOneResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateOneResponse;
  static deserializeBinaryFromReader(message: CreateOneResponse, reader: jspb.BinaryReader): CreateOneResponse;
}

export namespace CreateOneResponse {
  export type AsObject = {
    result?: protobuf_os_os_pb.OS.AsObject,
  }
}


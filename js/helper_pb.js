// source: helper.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

goog.exportSymbol('proto.simrs.Accounting', null, global);
goog.exportSymbol('proto.simrs.HelperVar', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.simrs.HelperVar = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.simrs.HelperVar, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.simrs.HelperVar.displayName = 'proto.simrs.HelperVar';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.simrs.Accounting = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.simrs.Accounting, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.simrs.Accounting.displayName = 'proto.simrs.Accounting';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.simrs.HelperVar.prototype.toObject = function(opt_includeInstance) {
  return proto.simrs.HelperVar.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.simrs.HelperVar} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.simrs.HelperVar.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.simrs.HelperVar}
 */
proto.simrs.HelperVar.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.simrs.HelperVar;
  return proto.simrs.HelperVar.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.simrs.HelperVar} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.simrs.HelperVar}
 */
proto.simrs.HelperVar.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.simrs.HelperVar.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.simrs.HelperVar.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.simrs.HelperVar} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.simrs.HelperVar.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional int64 id = 1;
 * @return {number}
 */
proto.simrs.HelperVar.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.simrs.HelperVar} returns this
 */
proto.simrs.HelperVar.prototype.setId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.simrs.HelperVar.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.simrs.HelperVar} returns this
 */
proto.simrs.HelperVar.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.simrs.Accounting.prototype.toObject = function(opt_includeInstance) {
  return proto.simrs.Accounting.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.simrs.Accounting} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.simrs.Accounting.toObject = function(includeInstance, msg) {
  var f, obj = {
    createby: jspb.Message.getFieldWithDefault(msg, 1, 0),
    createat: jspb.Message.getFieldWithDefault(msg, 2, ""),
    updateby: jspb.Message.getFieldWithDefault(msg, 3, 0),
    updateat: jspb.Message.getFieldWithDefault(msg, 4, ""),
    deleteby: jspb.Message.getFieldWithDefault(msg, 5, 0),
    deleteat: jspb.Message.getFieldWithDefault(msg, 6, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.simrs.Accounting}
 */
proto.simrs.Accounting.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.simrs.Accounting;
  return proto.simrs.Accounting.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.simrs.Accounting} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.simrs.Accounting}
 */
proto.simrs.Accounting.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCreateby(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setCreateat(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUpdateby(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setUpdateat(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setDeleteby(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setDeleteat(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.simrs.Accounting.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.simrs.Accounting.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.simrs.Accounting} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.simrs.Accounting.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCreateby();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getCreateat();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getUpdateby();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getUpdateat();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getDeleteby();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getDeleteat();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional int64 createBy = 1;
 * @return {number}
 */
proto.simrs.Accounting.prototype.getCreateby = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.simrs.Accounting} returns this
 */
proto.simrs.Accounting.prototype.setCreateby = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string createAt = 2;
 * @return {string}
 */
proto.simrs.Accounting.prototype.getCreateat = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.simrs.Accounting} returns this
 */
proto.simrs.Accounting.prototype.setCreateat = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int64 updateBy = 3;
 * @return {number}
 */
proto.simrs.Accounting.prototype.getUpdateby = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.simrs.Accounting} returns this
 */
proto.simrs.Accounting.prototype.setUpdateby = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional string updateAt = 4;
 * @return {string}
 */
proto.simrs.Accounting.prototype.getUpdateat = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.simrs.Accounting} returns this
 */
proto.simrs.Accounting.prototype.setUpdateat = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int64 deleteBy = 5;
 * @return {number}
 */
proto.simrs.Accounting.prototype.getDeleteby = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.simrs.Accounting} returns this
 */
proto.simrs.Accounting.prototype.setDeleteby = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional string deleteAt = 6;
 * @return {string}
 */
proto.simrs.Accounting.prototype.getDeleteat = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.simrs.Accounting} returns this
 */
proto.simrs.Accounting.prototype.setDeleteat = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


goog.object.extend(exports, proto.simrs);

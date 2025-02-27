/**
 * @fileoverview gRPC-Web generated client stub for simrs
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v3.21.12
// source: simrs.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var encounter_pb = require('./encounter_pb.js')

var patient_pb = require('./patient_pb.js')

var front_desk_pb = require('./front_desk_pb.js')

var polyclinic_pb = require('./polyclinic_pb.js')

var emergency_room_pb = require('./emergency_room_pb.js')
const proto = {};
proto.simrs = require('./simrs_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.PatientServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.PatientServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.GetPatientsRequest,
 *   !proto.simrs.GetPatientsResponse>}
 */
const methodDescriptor_PatientService_GetPatients = new grpc.web.MethodDescriptor(
  '/simrs.PatientService/GetPatients',
  grpc.web.MethodType.UNARY,
  patient_pb.GetPatientsRequest,
  patient_pb.GetPatientsResponse,
  /**
   * @param {!proto.simrs.GetPatientsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  patient_pb.GetPatientsResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.GetPatientsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.GetPatientsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.GetPatientsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.PatientServiceClient.prototype.getPatients =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.PatientService/GetPatients',
      request,
      metadata || {},
      methodDescriptor_PatientService_GetPatients,
      callback);
};


/**
 * @param {!proto.simrs.GetPatientsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.GetPatientsResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.PatientServicePromiseClient.prototype.getPatients =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.PatientService/GetPatients',
      request,
      metadata || {},
      methodDescriptor_PatientService_GetPatients);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.EncounterServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.EncounterServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.GetEncounterDetailsRequest,
 *   !proto.simrs.GetEncounterDetailsResponse>}
 */
const methodDescriptor_EncounterService_GetEncounterDetails = new grpc.web.MethodDescriptor(
  '/simrs.EncounterService/GetEncounterDetails',
  grpc.web.MethodType.UNARY,
  encounter_pb.GetEncounterDetailsRequest,
  encounter_pb.GetEncounterDetailsResponse,
  /**
   * @param {!proto.simrs.GetEncounterDetailsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  encounter_pb.GetEncounterDetailsResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.GetEncounterDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.GetEncounterDetailsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.GetEncounterDetailsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.EncounterServiceClient.prototype.getEncounterDetails =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.EncounterService/GetEncounterDetails',
      request,
      metadata || {},
      methodDescriptor_EncounterService_GetEncounterDetails,
      callback);
};


/**
 * @param {!proto.simrs.GetEncounterDetailsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.GetEncounterDetailsResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.EncounterServicePromiseClient.prototype.getEncounterDetails =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.EncounterService/GetEncounterDetails',
      request,
      metadata || {},
      methodDescriptor_EncounterService_GetEncounterDetails);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.AddEncounterRoomHistoryRequest,
 *   !proto.simrs.GenericEncounterServiceResponse>}
 */
const methodDescriptor_EncounterService_AddEncounterRoomHistory = new grpc.web.MethodDescriptor(
  '/simrs.EncounterService/AddEncounterRoomHistory',
  grpc.web.MethodType.UNARY,
  encounter_pb.AddEncounterRoomHistoryRequest,
  encounter_pb.GenericEncounterServiceResponse,
  /**
   * @param {!proto.simrs.AddEncounterRoomHistoryRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  encounter_pb.GenericEncounterServiceResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.AddEncounterRoomHistoryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.GenericEncounterServiceResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.GenericEncounterServiceResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.EncounterServiceClient.prototype.addEncounterRoomHistory =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.EncounterService/AddEncounterRoomHistory',
      request,
      metadata || {},
      methodDescriptor_EncounterService_AddEncounterRoomHistory,
      callback);
};


/**
 * @param {!proto.simrs.AddEncounterRoomHistoryRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.GenericEncounterServiceResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.EncounterServicePromiseClient.prototype.addEncounterRoomHistory =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.EncounterService/AddEncounterRoomHistory',
      request,
      metadata || {},
      methodDescriptor_EncounterService_AddEncounterRoomHistory);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.DeleteEncounterRoomHistoryRequest,
 *   !proto.simrs.GenericEncounterServiceResponse>}
 */
const methodDescriptor_EncounterService_DeleteEncounterRoomHistory = new grpc.web.MethodDescriptor(
  '/simrs.EncounterService/DeleteEncounterRoomHistory',
  grpc.web.MethodType.UNARY,
  encounter_pb.DeleteEncounterRoomHistoryRequest,
  encounter_pb.GenericEncounterServiceResponse,
  /**
   * @param {!proto.simrs.DeleteEncounterRoomHistoryRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  encounter_pb.GenericEncounterServiceResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.DeleteEncounterRoomHistoryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.GenericEncounterServiceResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.GenericEncounterServiceResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.EncounterServiceClient.prototype.deleteEncounterRoomHistory =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.EncounterService/DeleteEncounterRoomHistory',
      request,
      metadata || {},
      methodDescriptor_EncounterService_DeleteEncounterRoomHistory,
      callback);
};


/**
 * @param {!proto.simrs.DeleteEncounterRoomHistoryRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.GenericEncounterServiceResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.EncounterServicePromiseClient.prototype.deleteEncounterRoomHistory =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.EncounterService/DeleteEncounterRoomHistory',
      request,
      metadata || {},
      methodDescriptor_EncounterService_DeleteEncounterRoomHistory);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.QueueServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.QueueServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.UpdateQueueRequest,
 *   !proto.simrs.UpdateQueueResponse>}
 */
const methodDescriptor_QueueService_UpdateQueue = new grpc.web.MethodDescriptor(
  '/simrs.QueueService/UpdateQueue',
  grpc.web.MethodType.UNARY,
  front_desk_pb.UpdateQueueRequest,
  front_desk_pb.UpdateQueueResponse,
  /**
   * @param {!proto.simrs.UpdateQueueRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  front_desk_pb.UpdateQueueResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.UpdateQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.UpdateQueueResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.UpdateQueueResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.QueueServiceClient.prototype.updateQueue =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.QueueService/UpdateQueue',
      request,
      metadata || {},
      methodDescriptor_QueueService_UpdateQueue,
      callback);
};


/**
 * @param {!proto.simrs.UpdateQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.UpdateQueueResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.QueueServicePromiseClient.prototype.updateQueue =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.QueueService/UpdateQueue',
      request,
      metadata || {},
      methodDescriptor_QueueService_UpdateQueue);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.CompensateUpdateQueueEncounterRequest,
 *   !proto.simrs.CompensateUpdateQueueEncounterResponse>}
 */
const methodDescriptor_QueueService_CompensateUpdateQueueEncounter = new grpc.web.MethodDescriptor(
  '/simrs.QueueService/CompensateUpdateQueueEncounter',
  grpc.web.MethodType.UNARY,
  front_desk_pb.CompensateUpdateQueueEncounterRequest,
  front_desk_pb.CompensateUpdateQueueEncounterResponse,
  /**
   * @param {!proto.simrs.CompensateUpdateQueueEncounterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  front_desk_pb.CompensateUpdateQueueEncounterResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.CompensateUpdateQueueEncounterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.CompensateUpdateQueueEncounterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.CompensateUpdateQueueEncounterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.QueueServiceClient.prototype.compensateUpdateQueueEncounter =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.QueueService/CompensateUpdateQueueEncounter',
      request,
      metadata || {},
      methodDescriptor_QueueService_CompensateUpdateQueueEncounter,
      callback);
};


/**
 * @param {!proto.simrs.CompensateUpdateQueueEncounterRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.CompensateUpdateQueueEncounterResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.QueueServicePromiseClient.prototype.compensateUpdateQueueEncounter =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.QueueService/CompensateUpdateQueueEncounter',
      request,
      metadata || {},
      methodDescriptor_QueueService_CompensateUpdateQueueEncounter);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.GetFrontDeskQueueRequest,
 *   !proto.simrs.GetFrontDeskQueueResponse>}
 */
const methodDescriptor_QueueService_GetFrontDeskQueue = new grpc.web.MethodDescriptor(
  '/simrs.QueueService/GetFrontDeskQueue',
  grpc.web.MethodType.UNARY,
  front_desk_pb.GetFrontDeskQueueRequest,
  front_desk_pb.GetFrontDeskQueueResponse,
  /**
   * @param {!proto.simrs.GetFrontDeskQueueRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  front_desk_pb.GetFrontDeskQueueResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.GetFrontDeskQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.GetFrontDeskQueueResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.GetFrontDeskQueueResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.QueueServiceClient.prototype.getFrontDeskQueue =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.QueueService/GetFrontDeskQueue',
      request,
      metadata || {},
      methodDescriptor_QueueService_GetFrontDeskQueue,
      callback);
};


/**
 * @param {!proto.simrs.GetFrontDeskQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.GetFrontDeskQueueResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.QueueServicePromiseClient.prototype.getFrontDeskQueue =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.QueueService/GetFrontDeskQueue',
      request,
      metadata || {},
      methodDescriptor_QueueService_GetFrontDeskQueue);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.ManualQueueRequest,
 *   !proto.simrs.ManualQueueResponse>}
 */
const methodDescriptor_QueueService_ManualQueue = new grpc.web.MethodDescriptor(
  '/simrs.QueueService/ManualQueue',
  grpc.web.MethodType.UNARY,
  front_desk_pb.ManualQueueRequest,
  front_desk_pb.ManualQueueResponse,
  /**
   * @param {!proto.simrs.ManualQueueRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  front_desk_pb.ManualQueueResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.ManualQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.ManualQueueResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.ManualQueueResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.QueueServiceClient.prototype.manualQueue =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.QueueService/ManualQueue',
      request,
      metadata || {},
      methodDescriptor_QueueService_ManualQueue,
      callback);
};


/**
 * @param {!proto.simrs.ManualQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.ManualQueueResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.QueueServicePromiseClient.prototype.manualQueue =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.QueueService/ManualQueue',
      request,
      metadata || {},
      methodDescriptor_QueueService_ManualQueue);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.CompensateManualQueueRequest,
 *   !proto.simrs.CompensateManualQueueResponse>}
 */
const methodDescriptor_QueueService_CompensateManualQueue = new grpc.web.MethodDescriptor(
  '/simrs.QueueService/CompensateManualQueue',
  grpc.web.MethodType.UNARY,
  front_desk_pb.CompensateManualQueueRequest,
  front_desk_pb.CompensateManualQueueResponse,
  /**
   * @param {!proto.simrs.CompensateManualQueueRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  front_desk_pb.CompensateManualQueueResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.CompensateManualQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.CompensateManualQueueResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.CompensateManualQueueResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.QueueServiceClient.prototype.compensateManualQueue =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.QueueService/CompensateManualQueue',
      request,
      metadata || {},
      methodDescriptor_QueueService_CompensateManualQueue,
      callback);
};


/**
 * @param {!proto.simrs.CompensateManualQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.CompensateManualQueueResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.QueueServicePromiseClient.prototype.compensateManualQueue =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.QueueService/CompensateManualQueue',
      request,
      metadata || {},
      methodDescriptor_QueueService_CompensateManualQueue);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.PolyclinicServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.PolyclinicServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.CreatePolyclinicQueueRequest,
 *   !proto.simrs.CreatePolyclinicQueueResponse>}
 */
const methodDescriptor_PolyclinicService_CreatePolyclinicQueue = new grpc.web.MethodDescriptor(
  '/simrs.PolyclinicService/CreatePolyclinicQueue',
  grpc.web.MethodType.UNARY,
  polyclinic_pb.CreatePolyclinicQueueRequest,
  polyclinic_pb.CreatePolyclinicQueueResponse,
  /**
   * @param {!proto.simrs.CreatePolyclinicQueueRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  polyclinic_pb.CreatePolyclinicQueueResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.CreatePolyclinicQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.CreatePolyclinicQueueResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.CreatePolyclinicQueueResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.PolyclinicServiceClient.prototype.createPolyclinicQueue =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.PolyclinicService/CreatePolyclinicQueue',
      request,
      metadata || {},
      methodDescriptor_PolyclinicService_CreatePolyclinicQueue,
      callback);
};


/**
 * @param {!proto.simrs.CreatePolyclinicQueueRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.CreatePolyclinicQueueResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.PolyclinicServicePromiseClient.prototype.createPolyclinicQueue =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.PolyclinicService/CreatePolyclinicQueue',
      request,
      metadata || {},
      methodDescriptor_PolyclinicService_CreatePolyclinicQueue);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.CompensateCreatePolyQueueEncounterRequest,
 *   !proto.simrs.CompensateCreatePolyQueueEncounterResponse>}
 */
const methodDescriptor_PolyclinicService_CompensateCreatePolyQueueEncounter = new grpc.web.MethodDescriptor(
  '/simrs.PolyclinicService/CompensateCreatePolyQueueEncounter',
  grpc.web.MethodType.UNARY,
  polyclinic_pb.CompensateCreatePolyQueueEncounterRequest,
  polyclinic_pb.CompensateCreatePolyQueueEncounterResponse,
  /**
   * @param {!proto.simrs.CompensateCreatePolyQueueEncounterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  polyclinic_pb.CompensateCreatePolyQueueEncounterResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.CompensateCreatePolyQueueEncounterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.CompensateCreatePolyQueueEncounterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.CompensateCreatePolyQueueEncounterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.PolyclinicServiceClient.prototype.compensateCreatePolyQueueEncounter =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.PolyclinicService/CompensateCreatePolyQueueEncounter',
      request,
      metadata || {},
      methodDescriptor_PolyclinicService_CompensateCreatePolyQueueEncounter,
      callback);
};


/**
 * @param {!proto.simrs.CompensateCreatePolyQueueEncounterRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.CompensateCreatePolyQueueEncounterResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.PolyclinicServicePromiseClient.prototype.compensateCreatePolyQueueEncounter =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.PolyclinicService/CompensateCreatePolyQueueEncounter',
      request,
      metadata || {},
      methodDescriptor_PolyclinicService_CompensateCreatePolyQueueEncounter);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.EmergencyRoomServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simrs.EmergencyRoomServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.AddERPatientRequest,
 *   !proto.simrs.GenericERServiceResponse>}
 */
const methodDescriptor_EmergencyRoomService_AddERPatient = new grpc.web.MethodDescriptor(
  '/simrs.EmergencyRoomService/AddERPatient',
  grpc.web.MethodType.UNARY,
  emergency_room_pb.AddERPatientRequest,
  emergency_room_pb.GenericERServiceResponse,
  /**
   * @param {!proto.simrs.AddERPatientRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  emergency_room_pb.GenericERServiceResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.AddERPatientRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.GenericERServiceResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.GenericERServiceResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.EmergencyRoomServiceClient.prototype.addERPatient =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.EmergencyRoomService/AddERPatient',
      request,
      metadata || {},
      methodDescriptor_EmergencyRoomService_AddERPatient,
      callback);
};


/**
 * @param {!proto.simrs.AddERPatientRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.GenericERServiceResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.EmergencyRoomServicePromiseClient.prototype.addERPatient =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.EmergencyRoomService/AddERPatient',
      request,
      metadata || {},
      methodDescriptor_EmergencyRoomService_AddERPatient);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.simrs.CompensateAddERPatientEncounterRequest,
 *   !proto.simrs.GenericERServiceResponse>}
 */
const methodDescriptor_EmergencyRoomService_CompensateAddERPatientEncounter = new grpc.web.MethodDescriptor(
  '/simrs.EmergencyRoomService/CompensateAddERPatientEncounter',
  grpc.web.MethodType.UNARY,
  emergency_room_pb.CompensateAddERPatientEncounterRequest,
  emergency_room_pb.GenericERServiceResponse,
  /**
   * @param {!proto.simrs.CompensateAddERPatientEncounterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  emergency_room_pb.GenericERServiceResponse.deserializeBinary
);


/**
 * @param {!proto.simrs.CompensateAddERPatientEncounterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simrs.GenericERServiceResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simrs.GenericERServiceResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simrs.EmergencyRoomServiceClient.prototype.compensateAddERPatientEncounter =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simrs.EmergencyRoomService/CompensateAddERPatientEncounter',
      request,
      metadata || {},
      methodDescriptor_EmergencyRoomService_CompensateAddERPatientEncounter,
      callback);
};


/**
 * @param {!proto.simrs.CompensateAddERPatientEncounterRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simrs.GenericERServiceResponse>}
 *     Promise that resolves to the response
 */
proto.simrs.EmergencyRoomServicePromiseClient.prototype.compensateAddERPatientEncounter =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simrs.EmergencyRoomService/CompensateAddERPatientEncounter',
      request,
      metadata || {},
      methodDescriptor_EmergencyRoomService_CompensateAddERPatientEncounter);
};


module.exports = proto.simrs;


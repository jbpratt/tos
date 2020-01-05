/**
 * @fileoverview gRPC-Web generated client stub for tospb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.tospb = require('./tos_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tospb.MenuServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tospb.MenuServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tospb.Empty,
 *   !proto.tospb.Menu>}
 */
const methodDescriptor_MenuService_GetMenu = new grpc.web.MethodDescriptor(
  '/tospb.MenuService/GetMenu',
  grpc.web.MethodType.UNARY,
  proto.tospb.Empty,
  proto.tospb.Menu,
  /**
   * @param {!proto.tospb.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Menu.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tospb.Empty,
 *   !proto.tospb.Menu>}
 */
const methodInfo_MenuService_GetMenu = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tospb.Menu,
  /**
   * @param {!proto.tospb.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Menu.deserializeBinary
);


/**
 * @param {!proto.tospb.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tospb.Menu)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.Menu>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tospb.MenuServiceClient.prototype.getMenu =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tospb.MenuService/GetMenu',
      request,
      metadata || {},
      methodDescriptor_MenuService_GetMenu,
      callback);
};


/**
 * @param {!proto.tospb.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tospb.Menu>}
 *     A native promise that resolves to the response
 */
proto.tospb.MenuServicePromiseClient.prototype.getMenu =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tospb.MenuService/GetMenu',
      request,
      metadata || {},
      methodDescriptor_MenuService_GetMenu);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tospb.Item,
 *   !proto.tospb.CreateMenuItemResponse>}
 */
const methodDescriptor_MenuService_CreateMenuItem = new grpc.web.MethodDescriptor(
  '/tospb.MenuService/CreateMenuItem',
  grpc.web.MethodType.UNARY,
  proto.tospb.Item,
  proto.tospb.CreateMenuItemResponse,
  /**
   * @param {!proto.tospb.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.CreateMenuItemResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tospb.Item,
 *   !proto.tospb.CreateMenuItemResponse>}
 */
const methodInfo_MenuService_CreateMenuItem = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tospb.CreateMenuItemResponse,
  /**
   * @param {!proto.tospb.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.CreateMenuItemResponse.deserializeBinary
);


/**
 * @param {!proto.tospb.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tospb.CreateMenuItemResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.CreateMenuItemResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tospb.MenuServiceClient.prototype.createMenuItem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tospb.MenuService/CreateMenuItem',
      request,
      metadata || {},
      methodDescriptor_MenuService_CreateMenuItem,
      callback);
};


/**
 * @param {!proto.tospb.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tospb.CreateMenuItemResponse>}
 *     A native promise that resolves to the response
 */
proto.tospb.MenuServicePromiseClient.prototype.createMenuItem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tospb.MenuService/CreateMenuItem',
      request,
      metadata || {},
      methodDescriptor_MenuService_CreateMenuItem);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tospb.Item,
 *   !proto.tospb.Response>}
 */
const methodDescriptor_MenuService_UpdateMenuItem = new grpc.web.MethodDescriptor(
  '/tospb.MenuService/UpdateMenuItem',
  grpc.web.MethodType.UNARY,
  proto.tospb.Item,
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tospb.Item,
 *   !proto.tospb.Response>}
 */
const methodInfo_MenuService_UpdateMenuItem = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.Item} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @param {!proto.tospb.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tospb.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tospb.MenuServiceClient.prototype.updateMenuItem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tospb.MenuService/UpdateMenuItem',
      request,
      metadata || {},
      methodDescriptor_MenuService_UpdateMenuItem,
      callback);
};


/**
 * @param {!proto.tospb.Item} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tospb.Response>}
 *     A native promise that resolves to the response
 */
proto.tospb.MenuServicePromiseClient.prototype.updateMenuItem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tospb.MenuService/UpdateMenuItem',
      request,
      metadata || {},
      methodDescriptor_MenuService_UpdateMenuItem);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tospb.DeleteMenuItemRequest,
 *   !proto.tospb.Response>}
 */
const methodDescriptor_MenuService_DeleteMenuItem = new grpc.web.MethodDescriptor(
  '/tospb.MenuService/DeleteMenuItem',
  grpc.web.MethodType.UNARY,
  proto.tospb.DeleteMenuItemRequest,
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.DeleteMenuItemRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tospb.DeleteMenuItemRequest,
 *   !proto.tospb.Response>}
 */
const methodInfo_MenuService_DeleteMenuItem = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.DeleteMenuItemRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @param {!proto.tospb.DeleteMenuItemRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tospb.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tospb.MenuServiceClient.prototype.deleteMenuItem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tospb.MenuService/DeleteMenuItem',
      request,
      metadata || {},
      methodDescriptor_MenuService_DeleteMenuItem,
      callback);
};


/**
 * @param {!proto.tospb.DeleteMenuItemRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tospb.Response>}
 *     A native promise that resolves to the response
 */
proto.tospb.MenuServicePromiseClient.prototype.deleteMenuItem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tospb.MenuService/DeleteMenuItem',
      request,
      metadata || {},
      methodDescriptor_MenuService_DeleteMenuItem);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tospb.Option,
 *   !proto.tospb.Response>}
 */
const methodDescriptor_MenuService_CreateMenuItemOption = new grpc.web.MethodDescriptor(
  '/tospb.MenuService/CreateMenuItemOption',
  grpc.web.MethodType.UNARY,
  proto.tospb.Option,
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.Option} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tospb.Option,
 *   !proto.tospb.Response>}
 */
const methodInfo_MenuService_CreateMenuItemOption = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.Option} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @param {!proto.tospb.Option} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tospb.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tospb.MenuServiceClient.prototype.createMenuItemOption =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tospb.MenuService/CreateMenuItemOption',
      request,
      metadata || {},
      methodDescriptor_MenuService_CreateMenuItemOption,
      callback);
};


/**
 * @param {!proto.tospb.Option} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tospb.Response>}
 *     A native promise that resolves to the response
 */
proto.tospb.MenuServicePromiseClient.prototype.createMenuItemOption =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tospb.MenuService/CreateMenuItemOption',
      request,
      metadata || {},
      methodDescriptor_MenuService_CreateMenuItemOption);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tospb.OrderServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tospb.OrderServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tospb.Order,
 *   !proto.tospb.Response>}
 */
const methodDescriptor_OrderService_SubmitOrder = new grpc.web.MethodDescriptor(
  '/tospb.OrderService/SubmitOrder',
  grpc.web.MethodType.UNARY,
  proto.tospb.Order,
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tospb.Order,
 *   !proto.tospb.Response>}
 */
const methodInfo_OrderService_SubmitOrder = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @param {!proto.tospb.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tospb.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tospb.OrderServiceClient.prototype.submitOrder =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tospb.OrderService/SubmitOrder',
      request,
      metadata || {},
      methodDescriptor_OrderService_SubmitOrder,
      callback);
};


/**
 * @param {!proto.tospb.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tospb.Response>}
 *     A native promise that resolves to the response
 */
proto.tospb.OrderServicePromiseClient.prototype.submitOrder =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tospb.OrderService/SubmitOrder',
      request,
      metadata || {},
      methodDescriptor_OrderService_SubmitOrder);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tospb.Empty,
 *   !proto.tospb.OrdersResponse>}
 */
const methodDescriptor_OrderService_ActiveOrders = new grpc.web.MethodDescriptor(
  '/tospb.OrderService/ActiveOrders',
  grpc.web.MethodType.UNARY,
  proto.tospb.Empty,
  proto.tospb.OrdersResponse,
  /**
   * @param {!proto.tospb.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.OrdersResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tospb.Empty,
 *   !proto.tospb.OrdersResponse>}
 */
const methodInfo_OrderService_ActiveOrders = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tospb.OrdersResponse,
  /**
   * @param {!proto.tospb.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.OrdersResponse.deserializeBinary
);


/**
 * @param {!proto.tospb.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tospb.OrdersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.OrdersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tospb.OrderServiceClient.prototype.activeOrders =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tospb.OrderService/ActiveOrders',
      request,
      metadata || {},
      methodDescriptor_OrderService_ActiveOrders,
      callback);
};


/**
 * @param {!proto.tospb.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tospb.OrdersResponse>}
 *     A native promise that resolves to the response
 */
proto.tospb.OrderServicePromiseClient.prototype.activeOrders =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tospb.OrderService/ActiveOrders',
      request,
      metadata || {},
      methodDescriptor_OrderService_ActiveOrders);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tospb.CompleteOrderRequest,
 *   !proto.tospb.Response>}
 */
const methodDescriptor_OrderService_CompleteOrder = new grpc.web.MethodDescriptor(
  '/tospb.OrderService/CompleteOrder',
  grpc.web.MethodType.UNARY,
  proto.tospb.CompleteOrderRequest,
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.CompleteOrderRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tospb.CompleteOrderRequest,
 *   !proto.tospb.Response>}
 */
const methodInfo_OrderService_CompleteOrder = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tospb.Response,
  /**
   * @param {!proto.tospb.CompleteOrderRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Response.deserializeBinary
);


/**
 * @param {!proto.tospb.CompleteOrderRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tospb.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tospb.OrderServiceClient.prototype.completeOrder =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tospb.OrderService/CompleteOrder',
      request,
      metadata || {},
      methodDescriptor_OrderService_CompleteOrder,
      callback);
};


/**
 * @param {!proto.tospb.CompleteOrderRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tospb.Response>}
 *     A native promise that resolves to the response
 */
proto.tospb.OrderServicePromiseClient.prototype.completeOrder =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tospb.OrderService/CompleteOrder',
      request,
      metadata || {},
      methodDescriptor_OrderService_CompleteOrder);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tospb.Empty,
 *   !proto.tospb.Order>}
 */
const methodDescriptor_OrderService_SubscribeToOrders = new grpc.web.MethodDescriptor(
  '/tospb.OrderService/SubscribeToOrders',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.tospb.Empty,
  proto.tospb.Order,
  /**
   * @param {!proto.tospb.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Order.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tospb.Empty,
 *   !proto.tospb.Order>}
 */
const methodInfo_OrderService_SubscribeToOrders = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tospb.Order,
  /**
   * @param {!proto.tospb.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tospb.Order.deserializeBinary
);


/**
 * @param {!proto.tospb.Empty} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.Order>}
 *     The XHR Node Readable Stream
 */
proto.tospb.OrderServiceClient.prototype.subscribeToOrders =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/tospb.OrderService/SubscribeToOrders',
      request,
      metadata || {},
      methodDescriptor_OrderService_SubscribeToOrders);
};


/**
 * @param {!proto.tospb.Empty} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.tospb.Order>}
 *     The XHR Node Readable Stream
 */
proto.tospb.OrderServicePromiseClient.prototype.subscribeToOrders =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/tospb.OrderService/SubscribeToOrders',
      request,
      metadata || {},
      methodDescriptor_OrderService_SubscribeToOrders);
};


module.exports = proto.tospb;


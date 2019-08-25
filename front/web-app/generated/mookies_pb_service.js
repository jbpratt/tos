// package: mookiespb
// file: mookies.proto

var mookies_pb = require("./mookies_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var MenuService = (function () {
  function MenuService() {}
  MenuService.serviceName = "mookiespb.MenuService";
  return MenuService;
}());

MenuService.GetMenu = {
  methodName: "GetMenu",
  service: MenuService,
  requestStream: false,
  responseStream: false,
  requestType: mookies_pb.Empty,
  responseType: mookies_pb.Menu
};

MenuService.CreateMenuItem = {
  methodName: "CreateMenuItem",
  service: MenuService,
  requestStream: false,
  responseStream: false,
  requestType: mookies_pb.Item,
  responseType: mookies_pb.CreateMenuItemResponse
};

MenuService.UpdateMenuItem = {
  methodName: "UpdateMenuItem",
  service: MenuService,
  requestStream: false,
  responseStream: false,
  requestType: mookies_pb.Item,
  responseType: mookies_pb.Response
};

MenuService.DeleteMenuItem = {
  methodName: "DeleteMenuItem",
  service: MenuService,
  requestStream: false,
  responseStream: false,
  requestType: mookies_pb.DeleteMenuItemRequest,
  responseType: mookies_pb.Response
};

MenuService.CreateMenuItemOption = {
  methodName: "CreateMenuItemOption",
  service: MenuService,
  requestStream: false,
  responseStream: false,
  requestType: mookies_pb.Option,
  responseType: mookies_pb.Response
};

exports.MenuService = MenuService;

function MenuServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

MenuServiceClient.prototype.getMenu = function getMenu(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MenuService.GetMenu, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MenuServiceClient.prototype.createMenuItem = function createMenuItem(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MenuService.CreateMenuItem, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MenuServiceClient.prototype.updateMenuItem = function updateMenuItem(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MenuService.UpdateMenuItem, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MenuServiceClient.prototype.deleteMenuItem = function deleteMenuItem(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MenuService.DeleteMenuItem, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MenuServiceClient.prototype.createMenuItemOption = function createMenuItemOption(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MenuService.CreateMenuItemOption, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.MenuServiceClient = MenuServiceClient;

var OrderService = (function () {
  function OrderService() {}
  OrderService.serviceName = "mookiespb.OrderService";
  return OrderService;
}());

OrderService.SubmitOrder = {
  methodName: "SubmitOrder",
  service: OrderService,
  requestStream: false,
  responseStream: false,
  requestType: mookies_pb.Order,
  responseType: mookies_pb.Response
};

OrderService.ActiveOrders = {
  methodName: "ActiveOrders",
  service: OrderService,
  requestStream: false,
  responseStream: false,
  requestType: mookies_pb.Empty,
  responseType: mookies_pb.OrdersResponse
};

OrderService.CompleteOrder = {
  methodName: "CompleteOrder",
  service: OrderService,
  requestStream: false,
  responseStream: false,
  requestType: mookies_pb.CompleteOrderRequest,
  responseType: mookies_pb.Response
};

OrderService.SubscribeToOrders = {
  methodName: "SubscribeToOrders",
  service: OrderService,
  requestStream: false,
  responseStream: true,
  requestType: mookies_pb.Empty,
  responseType: mookies_pb.Order
};

exports.OrderService = OrderService;

function OrderServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

OrderServiceClient.prototype.submitOrder = function submitOrder(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderService.SubmitOrder, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

OrderServiceClient.prototype.activeOrders = function activeOrders(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderService.ActiveOrders, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

OrderServiceClient.prototype.completeOrder = function completeOrder(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderService.CompleteOrder, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

OrderServiceClient.prototype.subscribeToOrders = function subscribeToOrders(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(OrderService.SubscribeToOrders, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.OrderServiceClient = OrderServiceClient;


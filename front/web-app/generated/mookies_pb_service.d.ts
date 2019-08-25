// package: mookiespb
// file: mookies.proto

import * as mookies_pb from "./mookies_pb";
import {grpc} from "@improbable-eng/grpc-web";

type MenuServiceGetMenu = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mookies_pb.Empty;
  readonly responseType: typeof mookies_pb.Menu;
};

type MenuServiceCreateMenuItem = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mookies_pb.Item;
  readonly responseType: typeof mookies_pb.CreateMenuItemResponse;
};

type MenuServiceUpdateMenuItem = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mookies_pb.Item;
  readonly responseType: typeof mookies_pb.Response;
};

type MenuServiceDeleteMenuItem = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mookies_pb.DeleteMenuItemRequest;
  readonly responseType: typeof mookies_pb.Response;
};

type MenuServiceCreateMenuItemOption = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mookies_pb.Option;
  readonly responseType: typeof mookies_pb.Response;
};

export class MenuService {
  static readonly serviceName: string;
  static readonly GetMenu: MenuServiceGetMenu;
  static readonly CreateMenuItem: MenuServiceCreateMenuItem;
  static readonly UpdateMenuItem: MenuServiceUpdateMenuItem;
  static readonly DeleteMenuItem: MenuServiceDeleteMenuItem;
  static readonly CreateMenuItemOption: MenuServiceCreateMenuItemOption;
}

type OrderServiceSubmitOrder = {
  readonly methodName: string;
  readonly service: typeof OrderService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mookies_pb.Order;
  readonly responseType: typeof mookies_pb.Response;
};

type OrderServiceActiveOrders = {
  readonly methodName: string;
  readonly service: typeof OrderService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mookies_pb.Empty;
  readonly responseType: typeof mookies_pb.OrdersResponse;
};

type OrderServiceCompleteOrder = {
  readonly methodName: string;
  readonly service: typeof OrderService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof mookies_pb.CompleteOrderRequest;
  readonly responseType: typeof mookies_pb.Response;
};

type OrderServiceSubscribeToOrders = {
  readonly methodName: string;
  readonly service: typeof OrderService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof mookies_pb.Empty;
  readonly responseType: typeof mookies_pb.Order;
};

export class OrderService {
  static readonly serviceName: string;
  static readonly SubmitOrder: OrderServiceSubmitOrder;
  static readonly ActiveOrders: OrderServiceActiveOrders;
  static readonly CompleteOrder: OrderServiceCompleteOrder;
  static readonly SubscribeToOrders: OrderServiceSubscribeToOrders;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class MenuServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getMenu(
    requestMessage: mookies_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Menu|null) => void
  ): UnaryResponse;
  getMenu(
    requestMessage: mookies_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Menu|null) => void
  ): UnaryResponse;
  createMenuItem(
    requestMessage: mookies_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.CreateMenuItemResponse|null) => void
  ): UnaryResponse;
  createMenuItem(
    requestMessage: mookies_pb.Item,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.CreateMenuItemResponse|null) => void
  ): UnaryResponse;
  updateMenuItem(
    requestMessage: mookies_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
  updateMenuItem(
    requestMessage: mookies_pb.Item,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
  deleteMenuItem(
    requestMessage: mookies_pb.DeleteMenuItemRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
  deleteMenuItem(
    requestMessage: mookies_pb.DeleteMenuItemRequest,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
  createMenuItemOption(
    requestMessage: mookies_pb.Option,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
  createMenuItemOption(
    requestMessage: mookies_pb.Option,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
}

export class OrderServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  submitOrder(
    requestMessage: mookies_pb.Order,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
  submitOrder(
    requestMessage: mookies_pb.Order,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
  activeOrders(
    requestMessage: mookies_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.OrdersResponse|null) => void
  ): UnaryResponse;
  activeOrders(
    requestMessage: mookies_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.OrdersResponse|null) => void
  ): UnaryResponse;
  completeOrder(
    requestMessage: mookies_pb.CompleteOrderRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
  completeOrder(
    requestMessage: mookies_pb.CompleteOrderRequest,
    callback: (error: ServiceError|null, responseMessage: mookies_pb.Response|null) => void
  ): UnaryResponse;
  subscribeToOrders(requestMessage: mookies_pb.Empty, metadata?: grpc.Metadata): ResponseStream<mookies_pb.Order>;
}


// package: tospb
// file: protofiles/tos.proto

import * as protofiles_tos_pb from "../protofiles/tos_pb";
import {grpc} from "@improbable-eng/grpc-web";

type MenuServiceGetMenu = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof protofiles_tos_pb.Empty;
  readonly responseType: typeof protofiles_tos_pb.Menu;
};

type MenuServiceCreateMenuItem = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof protofiles_tos_pb.Item;
  readonly responseType: typeof protofiles_tos_pb.CreateMenuItemResponse;
};

type MenuServiceUpdateMenuItem = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof protofiles_tos_pb.Item;
  readonly responseType: typeof protofiles_tos_pb.Response;
};

type MenuServiceDeleteMenuItem = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof protofiles_tos_pb.DeleteMenuItemRequest;
  readonly responseType: typeof protofiles_tos_pb.Response;
};

type MenuServiceCreateMenuItemOption = {
  readonly methodName: string;
  readonly service: typeof MenuService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof protofiles_tos_pb.Option;
  readonly responseType: typeof protofiles_tos_pb.Response;
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
  readonly requestType: typeof protofiles_tos_pb.Order;
  readonly responseType: typeof protofiles_tos_pb.Response;
};

type OrderServiceActiveOrders = {
  readonly methodName: string;
  readonly service: typeof OrderService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof protofiles_tos_pb.Empty;
  readonly responseType: typeof protofiles_tos_pb.OrdersResponse;
};

type OrderServiceCompleteOrder = {
  readonly methodName: string;
  readonly service: typeof OrderService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof protofiles_tos_pb.CompleteOrderRequest;
  readonly responseType: typeof protofiles_tos_pb.Response;
};

type OrderServiceSubscribeToOrders = {
  readonly methodName: string;
  readonly service: typeof OrderService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof protofiles_tos_pb.Empty;
  readonly responseType: typeof protofiles_tos_pb.Order;
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
    requestMessage: protofiles_tos_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Menu|null) => void
  ): UnaryResponse;
  getMenu(
    requestMessage: protofiles_tos_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Menu|null) => void
  ): UnaryResponse;
  createMenuItem(
    requestMessage: protofiles_tos_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.CreateMenuItemResponse|null) => void
  ): UnaryResponse;
  createMenuItem(
    requestMessage: protofiles_tos_pb.Item,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.CreateMenuItemResponse|null) => void
  ): UnaryResponse;
  updateMenuItem(
    requestMessage: protofiles_tos_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
  updateMenuItem(
    requestMessage: protofiles_tos_pb.Item,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
  deleteMenuItem(
    requestMessage: protofiles_tos_pb.DeleteMenuItemRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
  deleteMenuItem(
    requestMessage: protofiles_tos_pb.DeleteMenuItemRequest,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
  createMenuItemOption(
    requestMessage: protofiles_tos_pb.Option,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
  createMenuItemOption(
    requestMessage: protofiles_tos_pb.Option,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
}

export class OrderServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  submitOrder(
    requestMessage: protofiles_tos_pb.Order,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
  submitOrder(
    requestMessage: protofiles_tos_pb.Order,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
  activeOrders(
    requestMessage: protofiles_tos_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.OrdersResponse|null) => void
  ): UnaryResponse;
  activeOrders(
    requestMessage: protofiles_tos_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.OrdersResponse|null) => void
  ): UnaryResponse;
  completeOrder(
    requestMessage: protofiles_tos_pb.CompleteOrderRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
  completeOrder(
    requestMessage: protofiles_tos_pb.CompleteOrderRequest,
    callback: (error: ServiceError|null, responseMessage: protofiles_tos_pb.Response|null) => void
  ): UnaryResponse;
  subscribeToOrders(requestMessage: protofiles_tos_pb.Empty, metadata?: grpc.Metadata): ResponseStream<protofiles_tos_pb.Order>;
}


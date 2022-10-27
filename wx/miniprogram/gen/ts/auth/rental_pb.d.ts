import * as $protobuf from "protobufjs";
/** Namespace rental. */
export namespace rental {

    /** Namespace v1. */
    namespace v1 {

        /** Properties of a Location. */
        interface ILocation {

            /** Location latitude */
            latitude?: (number | null);

            /** Location longitude */
            longitude?: (number | null);
        }

        /** Represents a Location. */
        class Location implements ILocation {

            /**
             * Constructs a new Location.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ILocation);

            /** Location latitude. */
            public latitude: number;

            /** Location longitude. */
            public longitude: number;

            /**
             * Decodes a Location message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Location
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.Location;

            /**
             * Creates a Location message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns Location
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.Location;

            /**
             * Creates a plain object from a Location message. Also converts values to other types if specified.
             * @param message Location
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.Location, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Location to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for Location
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a LocationStatus. */
        interface ILocationStatus {

            /** LocationStatus location */
            location?: (rental.v1.ILocation | null);

            /** LocationStatus feeCent */
            feeCent?: (number | null);

            /** LocationStatus kmDriven */
            kmDriven?: (number | null);

            /** LocationStatus poiName */
            poiName?: (string | null);

            /** LocationStatus timestampSec */
            timestampSec?: (number | $protobuf.Long | null);
        }

        /** Represents a LocationStatus. */
        class LocationStatus implements ILocationStatus {

            /**
             * Constructs a new LocationStatus.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ILocationStatus);

            /** LocationStatus location. */
            public location?: (rental.v1.ILocation | null);

            /** LocationStatus feeCent. */
            public feeCent: number;

            /** LocationStatus kmDriven. */
            public kmDriven: number;

            /** LocationStatus poiName. */
            public poiName: string;

            /** LocationStatus timestampSec. */
            public timestampSec: (number | $protobuf.Long);

            /**
             * Decodes a LocationStatus message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns LocationStatus
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.LocationStatus;

            /**
             * Creates a LocationStatus message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns LocationStatus
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.LocationStatus;

            /**
             * Creates a plain object from a LocationStatus message. Also converts values to other types if specified.
             * @param message LocationStatus
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.LocationStatus, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this LocationStatus to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for LocationStatus
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** TripStatus enum. */
        enum TripStatus {
            TS_NOT_SPECIFIED = 0,
            IN_PROGRESS = 1,
            FINISHED = 2
        }

        /** Properties of a TripEntity. */
        interface ITripEntity {

            /** TripEntity id */
            id?: (string | null);

            /** TripEntity trip */
            trip?: (rental.v1.ITrip | null);
        }

        /** Represents a TripEntity. */
        class TripEntity implements ITripEntity {

            /**
             * Constructs a new TripEntity.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ITripEntity);

            /** TripEntity id. */
            public id: string;

            /** TripEntity trip. */
            public trip?: (rental.v1.ITrip | null);

            /**
             * Decodes a TripEntity message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns TripEntity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.TripEntity;

            /**
             * Creates a TripEntity message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns TripEntity
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.TripEntity;

            /**
             * Creates a plain object from a TripEntity message. Also converts values to other types if specified.
             * @param message TripEntity
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.TripEntity, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this TripEntity to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for TripEntity
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a Trip. */
        interface ITrip {

            /** Trip accountId */
            accountId?: (string | null);

            /** Trip carId */
            carId?: (string | null);

            /** Trip start */
            start?: (rental.v1.ILocationStatus | null);

            /** Trip current */
            current?: (rental.v1.ILocationStatus | null);

            /** Trip end */
            end?: (rental.v1.ILocationStatus | null);

            /** Trip status */
            status?: (rental.v1.TripStatus | null);

            /** Trip identityId */
            identityId?: (string | null);
        }

        /** Represents a Trip. */
        class Trip implements ITrip {

            /**
             * Constructs a new Trip.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ITrip);

            /** Trip accountId. */
            public accountId: string;

            /** Trip carId. */
            public carId: string;

            /** Trip start. */
            public start?: (rental.v1.ILocationStatus | null);

            /** Trip current. */
            public current?: (rental.v1.ILocationStatus | null);

            /** Trip end. */
            public end?: (rental.v1.ILocationStatus | null);

            /** Trip status. */
            public status: rental.v1.TripStatus;

            /** Trip identityId. */
            public identityId: string;

            /**
             * Decodes a Trip message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Trip
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.Trip;

            /**
             * Creates a Trip message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns Trip
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.Trip;

            /**
             * Creates a plain object from a Trip message. Also converts values to other types if specified.
             * @param message Trip
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.Trip, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Trip to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for Trip
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a CreateTripRequest. */
        interface ICreateTripRequest {

            /** CreateTripRequest start */
            start?: (rental.v1.ILocation | null);

            /** CreateTripRequest carId */
            carId?: (string | null);
        }

        /** Represents a CreateTripRequest. */
        class CreateTripRequest implements ICreateTripRequest {

            /**
             * Constructs a new CreateTripRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ICreateTripRequest);

            /** CreateTripRequest start. */
            public start?: (rental.v1.ILocation | null);

            /** CreateTripRequest carId. */
            public carId: string;

            /**
             * Decodes a CreateTripRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns CreateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.CreateTripRequest;

            /**
             * Creates a CreateTripRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CreateTripRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.CreateTripRequest;

            /**
             * Creates a plain object from a CreateTripRequest message. Also converts values to other types if specified.
             * @param message CreateTripRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.CreateTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CreateTripRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for CreateTripRequest
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a GetTripRequest. */
        interface IGetTripRequest {

            /** GetTripRequest id */
            id?: (string | null);
        }

        /** Represents a GetTripRequest. */
        class GetTripRequest implements IGetTripRequest {

            /**
             * Constructs a new GetTripRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetTripRequest);

            /** GetTripRequest id. */
            public id: string;

            /**
             * Decodes a GetTripRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.GetTripRequest;

            /**
             * Creates a GetTripRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetTripRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetTripRequest;

            /**
             * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
             * @param message GetTripRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetTripRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for GetTripRequest
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a GetTripsRequest. */
        interface IGetTripsRequest {

            /** GetTripsRequest status */
            status?: (rental.v1.TripStatus | null);
        }

        /** Represents a GetTripsRequest. */
        class GetTripsRequest implements IGetTripsRequest {

            /**
             * Constructs a new GetTripsRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetTripsRequest);

            /** GetTripsRequest status. */
            public status: rental.v1.TripStatus;

            /**
             * Decodes a GetTripsRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetTripsRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.GetTripsRequest;

            /**
             * Creates a GetTripsRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetTripsRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetTripsRequest;

            /**
             * Creates a plain object from a GetTripsRequest message. Also converts values to other types if specified.
             * @param message GetTripsRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetTripsRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetTripsRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for GetTripsRequest
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a GetTripsResponse. */
        interface IGetTripsResponse {

            /** GetTripsResponse trips */
            trips?: (rental.v1.ITripEntity[] | null);
        }

        /** Represents a GetTripsResponse. */
        class GetTripsResponse implements IGetTripsResponse {

            /**
             * Constructs a new GetTripsResponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetTripsResponse);

            /** GetTripsResponse trips. */
            public trips: rental.v1.ITripEntity[];

            /**
             * Decodes a GetTripsResponse message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetTripsResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.GetTripsResponse;

            /**
             * Creates a GetTripsResponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetTripsResponse
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetTripsResponse;

            /**
             * Creates a plain object from a GetTripsResponse message. Also converts values to other types if specified.
             * @param message GetTripsResponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetTripsResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetTripsResponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for GetTripsResponse
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of an UpdateTripRequest. */
        interface IUpdateTripRequest {

            /** UpdateTripRequest id */
            id?: (string | null);

            /** UpdateTripRequest current */
            current?: (rental.v1.ILocation | null);

            /** UpdateTripRequest endTrip */
            endTrip?: (boolean | null);
        }

        /** Represents an UpdateTripRequest. */
        class UpdateTripRequest implements IUpdateTripRequest {

            /**
             * Constructs a new UpdateTripRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IUpdateTripRequest);

            /** UpdateTripRequest id. */
            public id: string;

            /** UpdateTripRequest current. */
            public current?: (rental.v1.ILocation | null);

            /** UpdateTripRequest endTrip. */
            public endTrip: boolean;

            /**
             * Decodes an UpdateTripRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns UpdateTripRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.UpdateTripRequest;

            /**
             * Creates an UpdateTripRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns UpdateTripRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.UpdateTripRequest;

            /**
             * Creates a plain object from an UpdateTripRequest message. Also converts values to other types if specified.
             * @param message UpdateTripRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.UpdateTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this UpdateTripRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for UpdateTripRequest
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Represents a TripService */
        class TripService extends $protobuf.rpc.Service {

            /**
             * Constructs a new TripService service.
             * @param rpcImpl RPC implementation
             * @param [requestDelimited=false] Whether requests are length-delimited
             * @param [responseDelimited=false] Whether responses are length-delimited
             */
            constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

            /**
             * Calls CreateTrip.
             * @param request CreateTripRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and TripEntity
             */
            public createTrip(request: rental.v1.ICreateTripRequest, callback: rental.v1.TripService.CreateTripCallback): void;

            /**
             * Calls CreateTrip.
             * @param request CreateTripRequest message or plain object
             * @returns Promise
             */
            public createTrip(request: rental.v1.ICreateTripRequest): Promise<rental.v1.TripEntity>;

            /**
             * Calls GetTrip.
             * @param request GetTripRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Trip
             */
            public getTrip(request: rental.v1.IGetTripRequest, callback: rental.v1.TripService.GetTripCallback): void;

            /**
             * Calls GetTrip.
             * @param request GetTripRequest message or plain object
             * @returns Promise
             */
            public getTrip(request: rental.v1.IGetTripRequest): Promise<rental.v1.Trip>;

            /**
             * Calls GetTrips.
             * @param request GetTripsRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and GetTripsResponse
             */
            public getTrips(request: rental.v1.IGetTripsRequest, callback: rental.v1.TripService.GetTripsCallback): void;

            /**
             * Calls GetTrips.
             * @param request GetTripsRequest message or plain object
             * @returns Promise
             */
            public getTrips(request: rental.v1.IGetTripsRequest): Promise<rental.v1.GetTripsResponse>;

            /**
             * Calls UpdateTrip.
             * @param request UpdateTripRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Trip
             */
            public updateTrip(request: rental.v1.IUpdateTripRequest, callback: rental.v1.TripService.UpdateTripCallback): void;

            /**
             * Calls UpdateTrip.
             * @param request UpdateTripRequest message or plain object
             * @returns Promise
             */
            public updateTrip(request: rental.v1.IUpdateTripRequest): Promise<rental.v1.Trip>;
        }

        namespace TripService {

            /**
             * Callback as used by {@link rental.v1.TripService#createTrip}.
             * @param error Error, if any
             * @param [response] TripEntity
             */
            type CreateTripCallback = (error: (Error | null), response?: rental.v1.TripEntity) => void;

            /**
             * Callback as used by {@link rental.v1.TripService#getTrip}.
             * @param error Error, if any
             * @param [response] Trip
             */
            type GetTripCallback = (error: (Error | null), response?: rental.v1.Trip) => void;

            /**
             * Callback as used by {@link rental.v1.TripService#getTrips}.
             * @param error Error, if any
             * @param [response] GetTripsResponse
             */
            type GetTripsCallback = (error: (Error | null), response?: rental.v1.GetTripsResponse) => void;

            /**
             * Callback as used by {@link rental.v1.TripService#updateTrip}.
             * @param error Error, if any
             * @param [response] Trip
             */
            type UpdateTripCallback = (error: (Error | null), response?: rental.v1.Trip) => void;
        }

        /** Gender enum. */
        enum Gender {
            G_NOT_SPECIFIED = 0,
            MALE = 1,
            FEMALE = 2
        }

        /** IdentityStatus enum. */
        enum IdentityStatus {
            UNSUBMITTED = 0,
            PENDING = 1,
            VERIFIED = 2
        }

        /** Properties of a Profile. */
        interface IProfile {

            /** Profile identity */
            identity?: (rental.v1.IIdentity | null);

            /** Profile identityStatus */
            identityStatus?: (rental.v1.IdentityStatus | null);
        }

        /** Represents a Profile. */
        class Profile implements IProfile {

            /**
             * Constructs a new Profile.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IProfile);

            /** Profile identity. */
            public identity?: (rental.v1.IIdentity | null);

            /** Profile identityStatus. */
            public identityStatus: rental.v1.IdentityStatus;

            /**
             * Decodes a Profile message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Profile
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.Profile;

            /**
             * Creates a Profile message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns Profile
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.Profile;

            /**
             * Creates a plain object from a Profile message. Also converts values to other types if specified.
             * @param message Profile
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.Profile, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Profile to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for Profile
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of an Identity. */
        interface IIdentity {

            /** Identity licNumber */
            licNumber?: (string | null);

            /** Identity name */
            name?: (string | null);

            /** Identity gender */
            gender?: (rental.v1.Gender | null);

            /** Identity birthDateMillis */
            birthDateMillis?: (number | $protobuf.Long | null);
        }

        /** Represents an Identity. */
        class Identity implements IIdentity {

            /**
             * Constructs a new Identity.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IIdentity);

            /** Identity licNumber. */
            public licNumber: string;

            /** Identity name. */
            public name: string;

            /** Identity gender. */
            public gender: rental.v1.Gender;

            /** Identity birthDateMillis. */
            public birthDateMillis: (number | $protobuf.Long);

            /**
             * Decodes an Identity message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Identity
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.Identity;

            /**
             * Creates an Identity message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns Identity
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.Identity;

            /**
             * Creates a plain object from an Identity message. Also converts values to other types if specified.
             * @param message Identity
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.Identity, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Identity to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for Identity
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a GetProfileRequest. */
        interface IGetProfileRequest {
        }

        /** Represents a GetProfileRequest. */
        class GetProfileRequest implements IGetProfileRequest {

            /**
             * Constructs a new GetProfileRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetProfileRequest);

            /**
             * Decodes a GetProfileRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.GetProfileRequest;

            /**
             * Creates a GetProfileRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetProfileRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetProfileRequest;

            /**
             * Creates a plain object from a GetProfileRequest message. Also converts values to other types if specified.
             * @param message GetProfileRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetProfileRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetProfileRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for GetProfileRequest
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a ClearProfileRequest. */
        interface IClearProfileRequest {
        }

        /** Represents a ClearProfileRequest. */
        class ClearProfileRequest implements IClearProfileRequest {

            /**
             * Constructs a new ClearProfileRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IClearProfileRequest);

            /**
             * Decodes a ClearProfileRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns ClearProfileRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.ClearProfileRequest;

            /**
             * Creates a ClearProfileRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns ClearProfileRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.ClearProfileRequest;

            /**
             * Creates a plain object from a ClearProfileRequest message. Also converts values to other types if specified.
             * @param message ClearProfileRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.ClearProfileRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this ClearProfileRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for ClearProfileRequest
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a GetProfilePhotoRequest. */
        interface IGetProfilePhotoRequest {
        }

        /** Represents a GetProfilePhotoRequest. */
        class GetProfilePhotoRequest implements IGetProfilePhotoRequest {

            /**
             * Constructs a new GetProfilePhotoRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetProfilePhotoRequest);

            /**
             * Decodes a GetProfilePhotoRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.GetProfilePhotoRequest;

            /**
             * Creates a GetProfilePhotoRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetProfilePhotoRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetProfilePhotoRequest;

            /**
             * Creates a plain object from a GetProfilePhotoRequest message. Also converts values to other types if specified.
             * @param message GetProfilePhotoRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetProfilePhotoRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetProfilePhotoRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for GetProfilePhotoRequest
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a GetProfilePhotoResponse. */
        interface IGetProfilePhotoResponse {

            /** GetProfilePhotoResponse uploadUrl */
            uploadUrl?: (string | null);
        }

        /** Represents a GetProfilePhotoResponse. */
        class GetProfilePhotoResponse implements IGetProfilePhotoResponse {

            /**
             * Constructs a new GetProfilePhotoResponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.IGetProfilePhotoResponse);

            /** GetProfilePhotoResponse uploadUrl. */
            public uploadUrl: string;

            /**
             * Decodes a GetProfilePhotoResponse message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GetProfilePhotoResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.GetProfilePhotoResponse;

            /**
             * Creates a GetProfilePhotoResponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns GetProfilePhotoResponse
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.GetProfilePhotoResponse;

            /**
             * Creates a plain object from a GetProfilePhotoResponse message. Also converts values to other types if specified.
             * @param message GetProfilePhotoResponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.GetProfilePhotoResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this GetProfilePhotoResponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for GetProfilePhotoResponse
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a CreateProfilePhotoRequest. */
        interface ICreateProfilePhotoRequest {
        }

        /** Represents a CreateProfilePhotoRequest. */
        class CreateProfilePhotoRequest implements ICreateProfilePhotoRequest {

            /**
             * Constructs a new CreateProfilePhotoRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ICreateProfilePhotoRequest);

            /**
             * Decodes a CreateProfilePhotoRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns CreateProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.CreateProfilePhotoRequest;

            /**
             * Creates a CreateProfilePhotoRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CreateProfilePhotoRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.CreateProfilePhotoRequest;

            /**
             * Creates a plain object from a CreateProfilePhotoRequest message. Also converts values to other types if specified.
             * @param message CreateProfilePhotoRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.CreateProfilePhotoRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CreateProfilePhotoRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for CreateProfilePhotoRequest
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a CreateProfilePhotoResponse. */
        interface ICreateProfilePhotoResponse {

            /** CreateProfilePhotoResponse uploadUrl */
            uploadUrl?: (string | null);
        }

        /** Represents a CreateProfilePhotoResponse. */
        class CreateProfilePhotoResponse implements ICreateProfilePhotoResponse {

            /**
             * Constructs a new CreateProfilePhotoResponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ICreateProfilePhotoResponse);

            /** CreateProfilePhotoResponse uploadUrl. */
            public uploadUrl: string;

            /**
             * Decodes a CreateProfilePhotoResponse message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns CreateProfilePhotoResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.CreateProfilePhotoResponse;

            /**
             * Creates a CreateProfilePhotoResponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CreateProfilePhotoResponse
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.CreateProfilePhotoResponse;

            /**
             * Creates a plain object from a CreateProfilePhotoResponse message. Also converts values to other types if specified.
             * @param message CreateProfilePhotoResponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.CreateProfilePhotoResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CreateProfilePhotoResponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for CreateProfilePhotoResponse
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Properties of a CompleteProfilePhotoRequest. */
        interface ICompleteProfilePhotoRequest {
        }

        /** Represents a CompleteProfilePhotoRequest. */
        class CompleteProfilePhotoRequest implements ICompleteProfilePhotoRequest {

            /**
             * Constructs a new CompleteProfilePhotoRequest.
             * @param [properties] Properties to set
             */
            constructor(properties?: rental.v1.ICompleteProfilePhotoRequest);

            /**
             * Decodes a CompleteProfilePhotoRequest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns CompleteProfilePhotoRequest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader | Uint8Array), length?: number): rental.v1.CompleteProfilePhotoRequest;

            /**
             * Creates a CompleteProfilePhotoRequest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns CompleteProfilePhotoRequest
             */
            public static fromObject(object: { [k: string]: any }): rental.v1.CompleteProfilePhotoRequest;

            /**
             * Creates a plain object from a CompleteProfilePhotoRequest message. Also converts values to other types if specified.
             * @param message CompleteProfilePhotoRequest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: rental.v1.CompleteProfilePhotoRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this CompleteProfilePhotoRequest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };

            /**
             * Gets the default type url for CompleteProfilePhotoRequest
             * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
             * @returns The default type url
             */
            public static getTypeUrl(typeUrlPrefix?: string): string;
        }

        /** Represents a ProfileService */
        class ProfileService extends $protobuf.rpc.Service {

            /**
             * Constructs a new ProfileService service.
             * @param rpcImpl RPC implementation
             * @param [requestDelimited=false] Whether requests are length-delimited
             * @param [responseDelimited=false] Whether responses are length-delimited
             */
            constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

            /**
             * Calls GetProfile.
             * @param request GetProfileRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Profile
             */
            public getProfile(request: rental.v1.IGetProfileRequest, callback: rental.v1.ProfileService.GetProfileCallback): void;

            /**
             * Calls GetProfile.
             * @param request GetProfileRequest message or plain object
             * @returns Promise
             */
            public getProfile(request: rental.v1.IGetProfileRequest): Promise<rental.v1.Profile>;

            /**
             * Calls SubmitProfile.
             * @param request Identity message or plain object
             * @param callback Node-style callback called with the error, if any, and Profile
             */
            public submitProfile(request: rental.v1.IIdentity, callback: rental.v1.ProfileService.SubmitProfileCallback): void;

            /**
             * Calls SubmitProfile.
             * @param request Identity message or plain object
             * @returns Promise
             */
            public submitProfile(request: rental.v1.IIdentity): Promise<rental.v1.Profile>;

            /**
             * Calls ClearProfile.
             * @param request ClearProfileRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Profile
             */
            public clearProfile(request: rental.v1.IClearProfileRequest, callback: rental.v1.ProfileService.ClearProfileCallback): void;

            /**
             * Calls ClearProfile.
             * @param request ClearProfileRequest message or plain object
             * @returns Promise
             */
            public clearProfile(request: rental.v1.IClearProfileRequest): Promise<rental.v1.Profile>;

            /**
             * Calls GetProfilePhoto.
             * @param request GetProfilePhotoRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and GetProfilePhotoResponse
             */
            public getProfilePhoto(request: rental.v1.IGetProfilePhotoRequest, callback: rental.v1.ProfileService.GetProfilePhotoCallback): void;

            /**
             * Calls GetProfilePhoto.
             * @param request GetProfilePhotoRequest message or plain object
             * @returns Promise
             */
            public getProfilePhoto(request: rental.v1.IGetProfilePhotoRequest): Promise<rental.v1.GetProfilePhotoResponse>;

            /**
             * Calls CreateProfilePhoto.
             * @param request CreateProfilePhotoRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and CreateProfilePhotoResponse
             */
            public createProfilePhoto(request: rental.v1.ICreateProfilePhotoRequest, callback: rental.v1.ProfileService.CreateProfilePhotoCallback): void;

            /**
             * Calls CreateProfilePhoto.
             * @param request CreateProfilePhotoRequest message or plain object
             * @returns Promise
             */
            public createProfilePhoto(request: rental.v1.ICreateProfilePhotoRequest): Promise<rental.v1.CreateProfilePhotoResponse>;

            /**
             * Calls CompleteProfilePhoto.
             * @param request CompleteProfilePhotoRequest message or plain object
             * @param callback Node-style callback called with the error, if any, and Identity
             */
            public completeProfilePhoto(request: rental.v1.ICompleteProfilePhotoRequest, callback: rental.v1.ProfileService.CompleteProfilePhotoCallback): void;

            /**
             * Calls CompleteProfilePhoto.
             * @param request CompleteProfilePhotoRequest message or plain object
             * @returns Promise
             */
            public completeProfilePhoto(request: rental.v1.ICompleteProfilePhotoRequest): Promise<rental.v1.Identity>;
        }

        namespace ProfileService {

            /**
             * Callback as used by {@link rental.v1.ProfileService#getProfile}.
             * @param error Error, if any
             * @param [response] Profile
             */
            type GetProfileCallback = (error: (Error | null), response?: rental.v1.Profile) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#submitProfile}.
             * @param error Error, if any
             * @param [response] Profile
             */
            type SubmitProfileCallback = (error: (Error | null), response?: rental.v1.Profile) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#clearProfile}.
             * @param error Error, if any
             * @param [response] Profile
             */
            type ClearProfileCallback = (error: (Error | null), response?: rental.v1.Profile) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#getProfilePhoto}.
             * @param error Error, if any
             * @param [response] GetProfilePhotoResponse
             */
            type GetProfilePhotoCallback = (error: (Error | null), response?: rental.v1.GetProfilePhotoResponse) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#createProfilePhoto}.
             * @param error Error, if any
             * @param [response] CreateProfilePhotoResponse
             */
            type CreateProfilePhotoCallback = (error: (Error | null), response?: rental.v1.CreateProfilePhotoResponse) => void;

            /**
             * Callback as used by {@link rental.v1.ProfileService#completeProfilePhoto}.
             * @param error Error, if any
             * @param [response] Identity
             */
            type CompleteProfilePhotoCallback = (error: (Error | null), response?: rental.v1.Identity) => void;
        }
    }
}

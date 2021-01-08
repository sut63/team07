/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntAmbulance,
    EntAmbulanceFromJSON,
    EntAmbulanceFromJSONTyped,
    EntAmbulanceToJSON,
    EntReceive,
    EntReceiveFromJSON,
    EntReceiveFromJSONTyped,
    EntReceiveToJSON,
    EntSend,
    EntSendFromJSON,
    EntSendFromJSONTyped,
    EntSendToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTransportEdges
 */
export interface EntTransportEdges {
    /**
     * 
     * @type {EntAmbulance}
     * @memberof EntTransportEdges
     */
    ambulance?: EntAmbulance;
    /**
     * 
     * @type {EntReceive}
     * @memberof EntTransportEdges
     */
    receiveid?: EntReceive;
    /**
     * 
     * @type {EntSend}
     * @memberof EntTransportEdges
     */
    sendid?: EntSend;
    /**
     * 
     * @type {EntUser}
     * @memberof EntTransportEdges
     */
    user?: EntUser;
}

export function EntTransportEdgesFromJSON(json: any): EntTransportEdges {
    return EntTransportEdgesFromJSONTyped(json, false);
}

export function EntTransportEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTransportEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ambulance': !exists(json, 'ambulance') ? undefined : EntAmbulanceFromJSON(json['ambulance']),
        'receiveid': !exists(json, 'receiveid') ? undefined : EntReceiveFromJSON(json['receiveid']),
        'sendid': !exists(json, 'sendid') ? undefined : EntSendFromJSON(json['sendid']),
        'user': !exists(json, 'user') ? undefined : EntUserFromJSON(json['user']),
    };
}

export function EntTransportEdgesToJSON(value?: EntTransportEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ambulance': EntAmbulanceToJSON(value.ambulance),
        'receiveid': EntReceiveToJSON(value.receiveid),
        'sendid': EntSendToJSON(value.sendid),
        'user': EntUserToJSON(value.user),
    };
}



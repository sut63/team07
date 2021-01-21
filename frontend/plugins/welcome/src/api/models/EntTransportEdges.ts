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
    EntHospital,
    EntHospitalFromJSON,
    EntHospitalFromJSONTyped,
    EntHospitalToJSON,
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
     * @type {EntHospital}
     * @memberof EntTransportEdges
     */
    receive?: EntHospital;
    /**
     * 
     * @type {EntHospital}
     * @memberof EntTransportEdges
     */
    send?: EntHospital;
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
        
        'ambulance': !exists(json, 'Ambulance') ? undefined : EntAmbulanceFromJSON(json['Ambulance']),
        'receive': !exists(json, 'Receive') ? undefined : EntHospitalFromJSON(json['Receive']),
        'send': !exists(json, 'Send') ? undefined : EntHospitalFromJSON(json['Send']),
        'user': !exists(json, 'User') ? undefined : EntUserFromJSON(json['User']),
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
        'receive': EntHospitalToJSON(value.receive),
        'send': EntHospitalToJSON(value.send),
        'user': EntUserToJSON(value.user),
    };
}



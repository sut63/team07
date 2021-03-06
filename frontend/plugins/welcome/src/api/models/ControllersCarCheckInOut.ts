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
/**
 * 
 * @export
 * @interface ControllersCarCheckInOut
 */
export interface ControllersCarCheckInOut {
    /**
     * 
     * @type {number}
     * @memberof ControllersCarCheckInOut
     */
    ambulance?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersCarCheckInOut
     */
    checkin?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersCarCheckInOut
     */
    checkout?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersCarCheckInOut
     */
    distance?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersCarCheckInOut
     */
    name?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersCarCheckInOut
     */
    note?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersCarCheckInOut
     */
    person?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersCarCheckInOut
     */
    place?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersCarCheckInOut
     */
    purpose?: number;
}

export function ControllersCarCheckInOutFromJSON(json: any): ControllersCarCheckInOut {
    return ControllersCarCheckInOutFromJSONTyped(json, false);
}

export function ControllersCarCheckInOutFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersCarCheckInOut {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ambulance': !exists(json, 'ambulance') ? undefined : json['ambulance'],
        'checkin': !exists(json, 'checkin') ? undefined : json['checkin'],
        'checkout': !exists(json, 'checkout') ? undefined : json['checkout'],
        'distance': !exists(json, 'distance') ? undefined : json['distance'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'note': !exists(json, 'note') ? undefined : json['note'],
        'person': !exists(json, 'person') ? undefined : json['person'],
        'place': !exists(json, 'place') ? undefined : json['place'],
        'purpose': !exists(json, 'purpose') ? undefined : json['purpose'],
    };
}

export function ControllersCarCheckInOutToJSON(value?: ControllersCarCheckInOut | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ambulance': value.ambulance,
        'checkin': value.checkin,
        'checkout': value.checkout,
        'distance': value.distance,
        'name': value.name,
        'note': value.note,
        'person': value.person,
        'place': value.place,
        'purpose': value.purpose,
    };
}



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
    EntCarCheckInOutEdges,
    EntCarCheckInOutEdgesFromJSON,
    EntCarCheckInOutEdgesFromJSONTyped,
    EntCarCheckInOutEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCarCheckInOut
 */
export interface EntCarCheckInOut {
    /**
     * CheckIn holds the value of the "checkIn" field.
     * @type {string}
     * @memberof EntCarCheckInOut
     */
    checkIn?: string;
    /**
     * CheckOut holds the value of the "checkOut" field.
     * @type {string}
     * @memberof EntCarCheckInOut
     */
    checkOut?: string;
    /**
     * 
     * @type {EntCarCheckInOutEdges}
     * @memberof EntCarCheckInOut
     */
    edges?: EntCarCheckInOutEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCarCheckInOut
     */
    id?: number;
    /**
     * Note holds the value of the "note" field.
     * @type {string}
     * @memberof EntCarCheckInOut
     */
    note?: string;
}

export function EntCarCheckInOutFromJSON(json: any): EntCarCheckInOut {
    return EntCarCheckInOutFromJSONTyped(json, false);
}

export function EntCarCheckInOutFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCarCheckInOut {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'checkIn': !exists(json, 'checkIn') ? undefined : json['checkIn'],
        'checkOut': !exists(json, 'checkOut') ? undefined : json['checkOut'],
        'edges': !exists(json, 'edges') ? undefined : EntCarCheckInOutEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'note': !exists(json, 'note') ? undefined : json['note'],
    };
}

export function EntCarCheckInOutToJSON(value?: EntCarCheckInOut | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'checkIn': value.checkIn,
        'checkOut': value.checkOut,
        'edges': EntCarCheckInOutEdgesToJSON(value.edges),
        'id': value.id,
        'note': value.note,
    };
}


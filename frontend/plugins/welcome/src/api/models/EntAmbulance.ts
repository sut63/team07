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
    EntAmbulanceEdges,
    EntAmbulanceEdgesFromJSON,
    EntAmbulanceEdgesFromJSONTyped,
    EntAmbulanceEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntAmbulance
 */
export interface EntAmbulance {
    /**
     * Carregistration holds the value of the "carregistration" field.
     * @type {string}
     * @memberof EntAmbulance
     */
    carregistration?: string;
    /**
     * 
     * @type {EntAmbulanceEdges}
     * @memberof EntAmbulance
     */
    edges?: EntAmbulanceEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntAmbulance
     */
    id?: number;
    /**
     * Registerat holds the value of the "registerat" field.
     * @type {string}
     * @memberof EntAmbulance
     */
    registerat?: string;
}

export function EntAmbulanceFromJSON(json: any): EntAmbulance {
    return EntAmbulanceFromJSONTyped(json, false);
}

export function EntAmbulanceFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntAmbulance {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'carregistration': !exists(json, 'carregistration') ? undefined : json['carregistration'],
        'edges': !exists(json, 'edges') ? undefined : EntAmbulanceEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'registerat': !exists(json, 'registerat') ? undefined : json['registerat'],
    };
}

export function EntAmbulanceToJSON(value?: EntAmbulance | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'carregistration': value.carregistration,
        'edges': EntAmbulanceEdgesToJSON(value.edges),
        'id': value.id,
        'registerat': value.registerat,
    };
}


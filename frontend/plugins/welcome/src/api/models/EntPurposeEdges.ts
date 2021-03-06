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
    EntCarCheckInOut,
    EntCarCheckInOutFromJSON,
    EntCarCheckInOutFromJSONTyped,
    EntCarCheckInOutToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPurposeEdges
 */
export interface EntPurposeEdges {
    /**
     * Carcheckinout holds the value of the carcheckinout edge.
     * @type {Array<EntCarCheckInOut>}
     * @memberof EntPurposeEdges
     */
    carcheckinout?: Array<EntCarCheckInOut>;
}

export function EntPurposeEdgesFromJSON(json: any): EntPurposeEdges {
    return EntPurposeEdgesFromJSONTyped(json, false);
}

export function EntPurposeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPurposeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'carcheckinout': !exists(json, 'Carcheckinout') ? undefined : ((json['Carcheckinout'] as Array<any>).map(EntCarCheckInOutFromJSON)),
    };
}

export function EntPurposeEdgesToJSON(value?: EntPurposeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'carcheckinout': value.carcheckinout === undefined ? undefined : ((value.carcheckinout as Array<any>).map(EntCarCheckInOutToJSON)),
    };
}



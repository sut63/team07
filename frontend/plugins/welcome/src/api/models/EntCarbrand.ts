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
    EntCarbrandEdges,
    EntCarbrandEdgesFromJSON,
    EntCarbrandEdgesFromJSONTyped,
    EntCarbrandEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCarbrand
 */
export interface EntCarbrand {
    /**
     * Brand holds the value of the "brand" field.
     * @type {string}
     * @memberof EntCarbrand
     */
    brand?: string;
    /**
     * 
     * @type {EntCarbrandEdges}
     * @memberof EntCarbrand
     */
    edges?: EntCarbrandEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCarbrand
     */
    id?: number;
}

export function EntCarbrandFromJSON(json: any): EntCarbrand {
    return EntCarbrandFromJSONTyped(json, false);
}

export function EntCarbrandFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCarbrand {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'brand': !exists(json, 'brand') ? undefined : json['brand'],
        'edges': !exists(json, 'edges') ? undefined : EntCarbrandEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntCarbrandToJSON(value?: EntCarbrand | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'brand': value.brand,
        'edges': EntCarbrandEdgesToJSON(value.edges),
        'id': value.id,
    };
}



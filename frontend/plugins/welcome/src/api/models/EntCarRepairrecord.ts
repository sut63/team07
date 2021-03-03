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
    EntCarRepairrecordEdges,
    EntCarRepairrecordEdgesFromJSON,
    EntCarRepairrecordEdgesFromJSONTyped,
    EntCarRepairrecordEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCarRepairrecord
 */
export interface EntCarRepairrecord {
    /**
     * Carmaintenance holds the value of the "carmaintenance" field.
     * @type {string}
     * @memberof EntCarRepairrecord
     */
    carmaintenance?: string;
    /**
     * Datetime holds the value of the "datetime" field.
     * @type {string}
     * @memberof EntCarRepairrecord
     */
    datetime?: string;
    /**
     * 
     * @type {EntCarRepairrecordEdges}
     * @memberof EntCarRepairrecord
     */
    edges?: EntCarRepairrecordEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCarRepairrecord
     */
    id?: number;
    /**
     * Repaircost holds the value of the "repaircost" field.
     * @type {number}
     * @memberof EntCarRepairrecord
     */
    repaircost?: number;
    /**
     * Repairdetail holds the value of the "repairdetail" field.
     * @type {string}
     * @memberof EntCarRepairrecord
     */
    repairdetail?: string;
}

export function EntCarRepairrecordFromJSON(json: any): EntCarRepairrecord {
    return EntCarRepairrecordFromJSONTyped(json, false);
}

export function EntCarRepairrecordFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCarRepairrecord {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'carmaintenance': !exists(json, 'carmaintenance') ? undefined : json['carmaintenance'],
        'datetime': !exists(json, 'datetime') ? undefined : json['datetime'],
        'edges': !exists(json, 'edges') ? undefined : EntCarRepairrecordEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'repaircost': !exists(json, 'repaircost') ? undefined : json['repaircost'],
        'repairdetail': !exists(json, 'repairdetail') ? undefined : json['repairdetail'],
    };
}

export function EntCarRepairrecordToJSON(value?: EntCarRepairrecord | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'carmaintenance': value.carmaintenance,
        'datetime': value.datetime,
        'edges': EntCarRepairrecordEdgesToJSON(value.edges),
        'id': value.id,
        'repaircost': value.repaircost,
        'repairdetail': value.repairdetail,
    };
}



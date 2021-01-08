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
    EntCarRepairrecord,
    EntCarRepairrecordFromJSON,
    EntCarRepairrecordFromJSONTyped,
    EntCarRepairrecordToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRepairingEdges
 */
export interface EntRepairingEdges {
    /**
     * Repairs holds the value of the repairs edge.
     * @type {Array<EntCarRepairrecord>}
     * @memberof EntRepairingEdges
     */
    repairs?: Array<EntCarRepairrecord>;
}

export function EntRepairingEdgesFromJSON(json: any): EntRepairingEdges {
    return EntRepairingEdgesFromJSONTyped(json, false);
}

export function EntRepairingEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRepairingEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'repairs': !exists(json, 'repairs') ? undefined : ((json['repairs'] as Array<any>).map(EntCarRepairrecordFromJSON)),
    };
}

export function EntRepairingEdgesToJSON(value?: EntRepairingEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'repairs': value.repairs === undefined ? undefined : ((value.repairs as Array<any>).map(EntCarRepairrecordToJSON)),
    };
}



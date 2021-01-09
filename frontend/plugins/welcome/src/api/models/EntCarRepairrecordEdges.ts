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
    EntCarInspection,
    EntCarInspectionFromJSON,
    EntCarInspectionFromJSONTyped,
    EntCarInspectionToJSON,
    EntRepairing,
    EntRepairingFromJSON,
    EntRepairingFromJSONTyped,
    EntRepairingToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCarRepairrecordEdges
 */
export interface EntCarRepairrecordEdges {
    /**
     * 
     * @type {EntCarInspection}
     * @memberof EntCarRepairrecordEdges
     */
    carinspection?: EntCarInspection;
    /**
     * 
     * @type {EntRepairing}
     * @memberof EntCarRepairrecordEdges
     */
    keeper?: EntRepairing;
    /**
     * 
     * @type {EntUser}
     * @memberof EntCarRepairrecordEdges
     */
    user?: EntUser;
}

export function EntCarRepairrecordEdgesFromJSON(json: any): EntCarRepairrecordEdges {
    return EntCarRepairrecordEdgesFromJSONTyped(json, false);
}

export function EntCarRepairrecordEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCarRepairrecordEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'carinspection': !exists(json, 'Carinspection') ? undefined : EntCarInspectionFromJSON(json['Carinspection']),
        'keeper': !exists(json, 'Keeper') ? undefined : EntRepairingFromJSON(json['Keeper']),
        'user': !exists(json, 'User') ? undefined : EntUserFromJSON(json['User']),
    };
}

export function EntCarRepairrecordEdgesToJSON(value?: EntCarRepairrecordEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'carinspection': EntCarInspectionToJSON(value.carinspection),
        'keeper': EntRepairingToJSON(value.keeper),
        'user': EntUserToJSON(value.user),
    };
}


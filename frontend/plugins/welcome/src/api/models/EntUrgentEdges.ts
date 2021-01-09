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
    EntCarservice,
    EntCarserviceFromJSON,
    EntCarserviceFromJSONTyped,
    EntCarserviceToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUrgentEdges
 */
export interface EntUrgentEdges {
    /**
     * Urgentid holds the value of the urgentid edge.
     * @type {Array<EntCarservice>}
     * @memberof EntUrgentEdges
     */
    urgentid?: Array<EntCarservice>;
}

export function EntUrgentEdgesFromJSON(json: any): EntUrgentEdges {
    return EntUrgentEdgesFromJSONTyped(json, false);
}

export function EntUrgentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUrgentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'urgentid': !exists(json, 'Urgentid') ? undefined : ((json['Urgentid'] as Array<any>).map(EntCarserviceFromJSON)),
    };
}

export function EntUrgentEdgesToJSON(value?: EntUrgentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'urgentid': value.urgentid === undefined ? undefined : ((value.urgentid as Array<any>).map(EntCarserviceToJSON)),
    };
}


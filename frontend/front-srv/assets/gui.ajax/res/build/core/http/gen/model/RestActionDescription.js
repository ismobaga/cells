/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

/**
* The RestActionDescription model module.
* @module model/RestActionDescription
* @version 1.0
*/

var RestActionDescription = (function () {
    /**
    * Constructs a new <code>RestActionDescription</code>.
    * @alias module:model/RestActionDescription
    * @class
    */

    function RestActionDescription() {
        _classCallCheck(this, RestActionDescription);

        this.Name = undefined;
        this.Icon = undefined;
        this.Label = undefined;
        this.Description = undefined;
        this.SummaryTemplate = undefined;
        this.HasForm = undefined;
        this.FormModule = undefined;
        this.FormModuleProps = undefined;
        this.Category = undefined;
        this.Tint = undefined;
        this.InputDescription = undefined;
        this.OutputDescription = undefined;
    }

    /**
    * Constructs a <code>RestActionDescription</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/RestActionDescription} obj Optional instance to populate.
    * @return {module:model/RestActionDescription} The populated <code>RestActionDescription</code> instance.
    */

    RestActionDescription.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RestActionDescription();

            if (data.hasOwnProperty('Name')) {
                obj['Name'] = _ApiClient2['default'].convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('Icon')) {
                obj['Icon'] = _ApiClient2['default'].convertToType(data['Icon'], 'String');
            }
            if (data.hasOwnProperty('Label')) {
                obj['Label'] = _ApiClient2['default'].convertToType(data['Label'], 'String');
            }
            if (data.hasOwnProperty('Description')) {
                obj['Description'] = _ApiClient2['default'].convertToType(data['Description'], 'String');
            }
            if (data.hasOwnProperty('SummaryTemplate')) {
                obj['SummaryTemplate'] = _ApiClient2['default'].convertToType(data['SummaryTemplate'], 'String');
            }
            if (data.hasOwnProperty('HasForm')) {
                obj['HasForm'] = _ApiClient2['default'].convertToType(data['HasForm'], 'Boolean');
            }
            if (data.hasOwnProperty('FormModule')) {
                obj['FormModule'] = _ApiClient2['default'].convertToType(data['FormModule'], 'String');
            }
            if (data.hasOwnProperty('FormModuleProps')) {
                obj['FormModuleProps'] = _ApiClient2['default'].convertToType(data['FormModuleProps'], 'String');
            }
            if (data.hasOwnProperty('Category')) {
                obj['Category'] = _ApiClient2['default'].convertToType(data['Category'], 'String');
            }
            if (data.hasOwnProperty('Tint')) {
                obj['Tint'] = _ApiClient2['default'].convertToType(data['Tint'], 'String');
            }
            if (data.hasOwnProperty('InputDescription')) {
                obj['InputDescription'] = _ApiClient2['default'].convertToType(data['InputDescription'], 'String');
            }
            if (data.hasOwnProperty('OutputDescription')) {
                obj['OutputDescription'] = _ApiClient2['default'].convertToType(data['OutputDescription'], 'String');
            }
        }
        return obj;
    };

    /**
    * @member {String} Name
    */
    return RestActionDescription;
})();

exports['default'] = RestActionDescription;
module.exports = exports['default'];

/**
* @member {String} Icon
*/

/**
* @member {String} Label
*/

/**
* @member {String} Description
*/

/**
* @member {String} SummaryTemplate
*/

/**
* @member {Boolean} HasForm
*/

/**
* @member {String} FormModule
*/

/**
* @member {String} FormModuleProps
*/

/**
* @member {String} Category
*/

/**
* @member {String} Tint
*/

/**
* @member {String} InputDescription
*/

/**
* @member {String} OutputDescription
*/

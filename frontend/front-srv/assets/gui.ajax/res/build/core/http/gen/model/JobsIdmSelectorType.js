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

"use strict";

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

/**
* Enum class JobsIdmSelectorType.
* @enum {}
* @readonly
*/

var JobsIdmSelectorType = (function () {
    function JobsIdmSelectorType() {
        _classCallCheck(this, JobsIdmSelectorType);

        this.User = "User";
        this.Role = "Role";
        this.Workspace = "Workspace";
        this.Acl = "Acl";
    }

    /**
    * Returns a <code>JobsIdmSelectorType</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/JobsIdmSelectorType} The enum <code>JobsIdmSelectorType</code> value.
    */

    JobsIdmSelectorType.constructFromObject = function constructFromObject(object) {
        return object;
    };

    return JobsIdmSelectorType;
})();

exports["default"] = JobsIdmSelectorType;
module.exports = exports["default"];

/**
 * value: "User"
 * @const
 */

/**
 * value: "Role"
 * @const
 */

/**
 * value: "Workspace"
 * @const
 */

/**
 * value: "Acl"
 * @const
 */
